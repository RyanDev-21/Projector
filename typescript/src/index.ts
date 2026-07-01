import getOpts from "./opts";
import config from "./config";

let opts = getOpts();
console.log(opts);
console.log(config(opts));

