// SPDX-License-Identifier: GPL-3.0-only
pragma solidity >=0.6.0 <=0.8.9;

contract Betting {
    /* Define the Bet struct */
    struct Bet {
        bytes32 outcome; // the guessed outcome
        uint256 amount; // the bet amount
    }

    address public owner; // the contract owner
    address public oracle; // the oracle that will decide the outcome of the betting
    address[] public gamblers; // list of the gamblers addresses
    mapping(address => bool) public isGambler;

    /* Maps the gambler's addresses to their bets */
    mapping(address => Bet) public bets;

    bool public decisionMade;

    /* List of all winners (maps are not iterable in solidity)*/
    address[] public winners;
    /* Maps the winners to their prize amount */
    mapping(address => uint256) public wins;

    /* Keep track of all possible outcomes */
    bytes32[] public outcomes;
    /* Map the valid outcomes for his betting */
    mapping(bytes32 => bool) public validOutcomes;
    /* Keep track of the total bet amount in all outcomes */
    mapping(bytes32 => uint256) public outcomeBets;

    
    event BetMade(
        address indexed gambler,
        bytes32 indexed outcome,
        uint256 amount
    );
    
    event Winners(address[] wins, uint256 totalPrize);
    event OracleChanged(
        address indexed previousOracle,
        address indexed newOracle
    );
    event Withdrawn(address indexed gambler, uint256 amount);

    
    modifier onlyOwner() {
        require(msg.sender == owner, "sender isn't the owner");
        _;
    }

    modifier onlyOracle() {
        require(isOracle(msg.sender), "sender isn't the oracle");
        _;
    }

    modifier requireOracle() {
        require(oracle != address(0), "no oracle found");
        _;
    }

    modifier outcomeExists(bytes32 outcome) {
        require(validOutcomes[outcome], "outcome not registered");
        _;
    }

    modifier onlyWinners() {
        require(wins[msg.sender] > 0, "sender should be a winner or insufficient requested amount ");
        _;
    }


    constructor(bytes32[] memory initOutcomes) {
        // register at least 2 possible outcomes,
        // and define the contract owner.
        require(initOutcomes.length >= 2,'must register at least 2 outcomes');
        owner=msg.sender;
        setOutcomes(initOutcomes);

    }

    function setOutcomes(bytes32[] memory _outcomes) public onlyOwner  {
        for (uint256 i = 0; i < _outcomes.length; i++) {
            outcomes.push(_outcomes[i]);
            validOutcomes[_outcomes[i]] = true;
        }
    }

    /**
     * @notice This function allows owner to chooses their trusted Oracle.
     * @param newOracle The address of the new oracle.
     */
    function chooseOracle(address newOracle) public onlyOwner {
        require(!isGambler[newOracle],"the oracle cannot be a gambler");
        require(newOracle!=owner,"the owner cannot be an oracle");
        emit OracleChanged(oracle,newOracle);
        oracle=newOracle;
    }

    /**
     * @notice Make a bet.
     * @param outcome The hash of the outcome to bet on.
     */
    function makeBet(bytes32 outcome) public payable requireOracle {
      //gamblers make bets and are registered
        require(msg.sender!=owner,"the owner cannot bet");
        require(!isOracle(msg.sender),"the oracle of the betting cannot bet");
        require(!decisionMade, "cannot bet after decision was made");
        require(!isGambler[msg.sender],"each gambler can only bet once");
        require(validOutcomes[outcome], "outcome not registered");
        bets[msg.sender].outcome = outcome;
        bets[msg.sender].amount = msg.value;
        gamblers.push(msg.sender);
        outcomeBets[outcome] += msg.value;
        isGambler[msg.sender] = true;
        emit BetMade(msg.sender, outcome, msg.value);
    }

    /**
     * @notice Decide on an outcome.
     * @param decidedOutcome The chosen outcome.
     */
    function makeDecision(bytes32 decidedOutcome) public onlyOracle outcomeExists(decidedOutcome) {

        require(!decisionMade,'can make decision only once');
        require(validOutcomes[decidedOutcome], "outcome not registered");
        require(gamblers.length>0);
        require(outcomes.length>0);
        decisionMade=true;

        uint256 sumWinAmount=0;
        uint256 totalPrize=0;
        for(uint256 i = 0; i < gamblers.length; i++) {
            if (bets[gamblers[i]].outcome == decidedOutcome){
                winners.push(gamblers[i]);
                sumWinAmount+=bets[gamblers[i]].amount;
            }
            totalPrize+=bets[gamblers[i]].amount;
        }

        //winnes get reward
        //if no gambler wins, the oracle gets all fund
        if (sumWinAmount==0){
            winners.push(oracle);
            wins[oracle]=totalPrize;
        } else{
            for(uint256 i = 0; i < winners.length; i++){
                wins[winners[i]]=bets[winners[i]].amount*totalPrize/sumWinAmount;
            }
        }

        emit Winners(winners,totalPrize);
    }

    /**
     * @notice This function allows the winners to withdraw their
     * winnings safely (if they win something).
     * @param amount The amount to be withdrawn
     */
    function withdraw(uint256 amount) public onlyWinners{
        require(wins[msg.sender]>=amount,'insufficient requested amount');    
        wins[msg.sender]-= amount;
        payable(msg.sender).transfer(amount);
        emit Withdrawn(msg.sender,amount);
        
    }

    /**
     * @notice Reset the contract state
     */
    function contractReset() public onlyOwner {
        // reset the contract variables to the initial state to allow new bettings and outcomes.
        require(decisionMade,"cannot reset before decision");
        decisionMade=false;
        for (uint256 i = 0; i <outcomes.length; i++){
            delete validOutcomes[outcomes[i]];
        }
        delete outcomes;
        delete winners;
        delete gamblers;

        
    }

    /**
     * @dev This function allows anyone to check the amount
     * already betted per outcome.
     * @param outcome The hash of the outcome to be checked.
     * @return The amount betted for the given outcome.
     */
    function checkOutcome(bytes32 outcome) public view returns (uint256) {
        // Returns the current stake for the given outcome
        require(validOutcomes[outcome],'outcome not registered');
        return outcomeBets[outcome];
    }

    /**
     * @notice This function is similar to `checkOutcome` but
     * it receives the outcome as string.
     * @param outcomeString The string representation of the outcome
     * to be checked.
     * @dev It uses the `keccak256` hash function to get the hash of the `outcomeString`
     * @return The amount betted for the given outcome.
     */
    function checkOutcomeString(string memory outcomeString)
        public
        view
        returns (uint256)
    {
        // hash outcomestring using keccak256. Then checkOutcome
        bytes32 outcome = keccak256(bytes(outcomeString));
        return checkOutcome(outcome);
    }

    /**
     * @notice This function allows anyone to check their winnings.
     * @return The winning amount of the msg.sender if it exists.
     */
    function checkWinnings() public view returns (uint256) {
        /* TODO (students) */
        if (wins[msg.sender] > 0) {
            return wins[msg.sender];
        } else {
            return 0;
        }

    }

    function isOracle(address _oracle) public view returns (bool) {
        return oracle == _oracle;
    }

    function getGamblers() public view returns (address[] memory) {
        return gamblers;
    }

    function getWinners() public view returns (address[] memory) {
        return winners;
    }

    function getOutcomes() public view returns (bytes32[] memory) {
        return outcomes;
    }
}
