import { get } from "svelte/store";
import { token } from "stores/tokenStore";

class ModelBase {
  constructor(urlBase) {
    this.urlBase = urlBase;
  }

  getModel = async (id) => {
    const headers = new Headers({
      Token: get(token).token,
    });
    fetch(`${this.urlBase}/${id}`, { method: "GET", headers });
  };
}

export default ModelBase;
