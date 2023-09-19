/*
* This script is based on a JavaScript file found in the ZAP community-scripts GitHub repo:
* https://github.com/zaproxy/community-scripts/blob/main/httpsender/AddBearerTokenHeader.js
*/

function sendingRequest(msg, initiator, helper) {
    var token = org.zaproxy.zap.extension.script.ScriptVars.getGlobalVar("access_token")
    msg.getRequestHeader().setHeader("Authorization", "Bearer " + token);
    return msg;
  }
  
  function responseReceived(msg, initiator, helper) { }