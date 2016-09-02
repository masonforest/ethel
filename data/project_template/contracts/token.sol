contract Token {
  string public name;
  event Transfer(address indexed _from, address indexed _to, uint256 _value);
  mapping (address => uint) public balances;

  function Token() {
    balances[tx.origin] = 100;
  }

  function transfer(address receiver, uint amount) returns(bool sufficient) {
    if (balances[msg.sender] < amount) return false;
    balances[msg.sender] -= amount;
    balances[receiver] += amount;
    Transfer(msg.sender, receiver, amount);
    return true;
  }
}
