import { readable } from "svelte/store";
import Token from "models/token";

export const token = readable(null, (set) => {
  const tokenCookie = document.cookie
    .split(";")
    .find((cookiePair) => cookiePair.split("=")[0] === "token");

  if (!tokenCookie) {
    set(null);
    return;
  }

  const tokenJson = tokenCookie.split("=")[1];

  if (tokenJson) {
    const newUser = new Token(JSON.parse(tokenJson));
    set(newUser);
  }
});
