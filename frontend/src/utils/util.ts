
import * as wailsRuntime from "../../wailsjs/runtime/runtime";

export namespace Util {
  export function OpenUrl(url: string) {
    wailsRuntime.BrowserOpenURL(url)
  }
  
}