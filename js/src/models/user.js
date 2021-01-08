import ModelBase from "models/modelBase";

class User extends ModelBase {
  constructor(props) {
    super("/user");
    this.set(props);
  }

  id = null;
  fullName = null;
  username = null;

  set = ({ id, full_name, username } = {}) => {
    this.id = id;
    this.fullName = full_name;
    this.username = username;
  };
}

export default User;
