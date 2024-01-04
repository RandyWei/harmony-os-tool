
import * as wailsRuntime from "../../wailsjs/runtime/runtime";

export namespace Util {
  export function OpenUrl(url: string) {
    wailsRuntime.BrowserOpenURL(url)
  }
  export function Log(message:string) {
    wailsRuntime.LogInfo(message)
  }

  export function LogD(message:string) {
    wailsRuntime.LogDebug(message)
  }

  export function LogE(message:string) {
    wailsRuntime.LogError(message)
  }
}