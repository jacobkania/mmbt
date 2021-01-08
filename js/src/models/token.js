import ModelBase from "models/modelBase";

class Token {
  constructor(props) {
    this.set(props);
  }

  userID = null;
  token = null;

  set = ({ token, userID } = {}) => {
    this.userID = userID;
    this.token = token;
  };
}

export default Token;
