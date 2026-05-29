This file is a merged representation of the entire codebase, combined into a single document by Repomix.

<file_summary>
This section contains a summary of this file.

<purpose>
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.
</purpose>

<file_format>
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  - File path as an attribute
  - Full contents of the file
</file_format>

<usage_guidelines>
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.
</usage_guidelines>

<notes>
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)
</notes>

</file_summary>

<directory_structure>
api_server.py
decompiled.js
images/image-1.png
images/image-2.png
images/image-3.png
images/image-4.png
images/image.png
manual.py
readme.md
wrapper/__init__.py
wrapper/chatgpt.py
wrapper/IP_Info/headers.py
wrapper/IP_Info/ip_info.py
wrapper/logger.py
wrapper/reverse/challenges.py
wrapper/reverse/decompiler.py
wrapper/reverse/parse.py
wrapper/reverse/vm.py
wrapper/runtime.py
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path="api_server.py">
from fastapi      import FastAPI, HTTPException
from urllib.parse import urlparse, ParseResult
from pydantic     import BaseModel
from wrapper      import ChatGPT
from uvicorn      import run


app = FastAPI()

class ConversationRequest(BaseModel):
    proxy: str
    message: str
    image: str = None

def format_proxy(proxy: str) -> str:
    
    if not proxy.startswith(("http://", "https://")):
        proxy: str = "http://" + proxy
    
    try:
        parsed: ParseResult = urlparse(proxy)

        if parsed.scheme not in ("http", ""):
            raise ValueError("Not http scheme")

        if not parsed.hostname or not parsed.port:
            raise ValueError("No url and port")

        if parsed.username and parsed.password:
            return f"http://{parsed.username}:{parsed.password}@{parsed.hostname}:{parsed.port}"
        
        else:
            return f"http://{parsed.hostname}:{parsed.port}"
    
    except ValueError as e:
        raise HTTPException(status_code=400, detail=f"Invalid proxy format: {str(e)}")

@app.post("/conversation")
async def create_conversation(request: ConversationRequest):
    if not request.proxy or not request.message:
        raise HTTPException(status_code=400, detail="Proxy and message are required")
    
    proxy = format_proxy(request.proxy)
    
    try:
        if request.image:
            answer: str = ChatGPT(proxy).ask_question(request.message, request.image)
        else:
            answer: str = ChatGPT(proxy).ask_question(request.message)
        
        return {
            "status": "success",
            "result": answer
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error: {str(e)}")

if __name__ == "__main__":
    run(app, host="0.0.0.0", port=6969)
</file>

<file path="decompiled.js">
const { JSDOM } = require("jsdom");
const dom = new JSDOM("<!DOCTYPE html><p>Hello world</p>", { url: "https://chatgpt.com/" });
const window = dom.window;
var mem = {};
function XOR_STR(e, t) {
        e = String(e);
        t = String(t);
        let n = "";
        for (let r = 0; r < e.length; r++)
            n += String.fromCharCode(e.charCodeAt(r) ^ t.charCodeAt(r % t.length));
        return n;
    }
    
var var_10 = window;
var var_94_65 = 65.82;
var var_10_27 = "set";
var var_10_27 = window["Reflect"]["set"];
var var_17_41 = "create";
var var_17_41 = window["Object"]["create"];
var var_94_65 = XOR_STR(var_94_65, var_94_65);
var var_94_65 = btoa("" + var_94_65);
var var_86_99 = 86.06;
var var_36_48 = "localStorage";
var var_36_48 = window["localStorage"];
var var_61_37 = 2000;
var var_94_65 = XOR_STR(var_94_65, var_86_99);
var var_94_65 = btoa("" + var_94_65);
var var_33_97 = "now";
var var_33_97 = window["performance"]["now"].bind(window["performance"]);
var var_86_99 = 36.56;
var var_43_93 = var_33_97();
var var_52_53 = null;
var var_82_53 = var_17_41(var_52_53);
var var_86_99 = 50.43;
var var_88_1 = "navigator";
var var_67_47 = 29.45;
var var_88_1 = window["navigator"];
var var_20_04 = "aWIbBRYADwIRAAQXFwgUEgEbHwkLZAIRYwAKAAYJHhkWAxYHCgIRAAAXGwJlHhl1AwwcCRkdGAsNAAcNHhkXARYBCnMdGGkIFh8NAxUOCQocDB0dGAcMAAcLbxUOagsHFx0CFBIMGh8OARUOCQocDB1sZQ==";
var var_86_99 = 69.58;
var var_20_87 = 0.1;
var var_47_28 = 0.6000000000000001;
var var_55_63 = 48.51;
var var_78_68 = 19.33;
var var_20_87 = Array.isArray(var_20_87) ? (var_20_87.push(var_47_28), var_20_87) : var_20_87 + var_47_28;
var var_20_87 = XOR_STR(var_20_87, var_55_63);
var var_20_87 = btoa("" + var_20_87);
var_82_53[var_78_68] = var_20_87;
var var_33_8 = [];
var var_82_7 = "vendor";
try { mem[var_82_7] = var_88_1[var_82_7]; } catch(r) { var_33_8 = "" + r; }
var var_33_8 = Array.isArray(var_33_8) ? (var_33_8.push(var_82_7), var_33_8) : var_33_8 + var_82_7;
var var_34_13 = "platform";
try { mem[var_34_13] = var_88_1[var_34_13]; } catch(r) { var_33_8 = "" + r; }
var var_33_8 = Array.isArray(var_33_8) ? (var_33_8.push(var_34_13), var_33_8) : var_33_8 + var_34_13;
var var_2_52 = "deviceMemory";
var var_94_65 = 5.37;
try { mem[var_2_52] = var_88_1[var_2_52]; } catch(r) { var_33_8 = "" + r; }
var var_33_8 = Array.isArray(var_33_8) ? (var_33_8.push(var_2_52), var_33_8) : var_33_8 + var_2_52;
var var_50_72 = "maxTouchPoints";
try { mem[var_50_72] = var_88_1[var_50_72]; } catch(r) { var_33_8 = "" + r; }
var var_33_8 = Array.isArray(var_33_8) ? (var_33_8.push(var_50_72), var_33_8) : var_33_8 + var_50_72;
var var_33_8 = JSON.stringify(var_33_8);
var var_33_8 = XOR_STR(var_33_8, var_55_63);
var var_94_65 = XOR_STR(var_94_65, var_94_65);
var var_94_65 = btoa("" + var_94_65);
var var_33_8 = btoa("" + var_33_8);
var var_57_57 = 56.04;
var_82_53[var_57_57] = var_33_8;
var var_55_63 = 48.51;
var var_95_86 = 14.85;
var var_5_35 = "createElement";
var var_5_35 = window["document"]["createElement"].bind(window["document"]);
var var_12_62 = "div";
var var_12_62 = var_5_35(var_12_62);
var var_89_53 = "style";
var var_89_53 = var_12_62["style"];
var var_70_3 = "hidden";
var var_63_32 = "visibility";
var_89_53[var_63_32] = var_70_3;
var var_94_65 = 7.04;
var var_46_21 = "ariaHidden";
var var_94_65 = XOR_STR(var_94_65, var_94_65);
var var_94_65 = btoa("" + var_94_65);
var var_52_83 = "True";
var_12_62[var_46_21] = var_52_83;
var var_42_87 = "position";
var var_41_49 = "fixed";
var var_86_99 = 61.16;
var_89_53[var_42_87] = var_41_49;
var var_36_61 = "Impact";
var var_22_22 = "fontFamily";
var_89_53[var_22_22] = var_36_61;
var var_65_3 = "10px";
var var_2_62 = "fontSize";
var var_77_52 = 43.91;
var var_86_99 = btoa("" + var_86_99);
var_82_53[var_77_52] = var_86_99;
var_89_53[var_2_62] = var_65_3;
var var_34_26 = "q̡̨̩̂̎ś̜̑̈G̦̭̟";
var var_96_27 = "innerText";
var_12_62[var_96_27] = var_34_26;
var var_71_49 = "appendChild";
var var_71_49 = window["document"]["body"]["appendChild"].bind(window["document"]["body"]);
var var_86_99 = 93.87;
var_71_49(var_12_62);
var var_14_87 = "getBoundingClientRect";
var var_94_65 = 90.6;
var var_14_87 = var_12_62["getBoundingClientRect"].bind(var_12_62);
var var_98_31 = var_14_87();
var var_98_31 = JSON.stringify(var_98_31);
var var_98_31 = XOR_STR(var_98_31, var_55_63);
var var_98_31 = btoa("" + var_98_31);
var_82_53[var_95_86] = var_98_31;
var var_94_99 = "removeChild";
var var_94_99 = window["document"]["body"]["removeChild"].bind(window["document"]["body"]);
var_94_99(var_12_62);
var var_4_41 = 31.17;
var var_44_93 = "keys";
var var_44_93 = window["Object"]["keys"];
var var_86_99 = 76.2;
var var_30_99 = var_44_93(var_36_48);
var var_30_99 = XOR_STR(var_30_99, var_55_63);
var var_30_99 = btoa("" + var_30_99);
var_82_53[var_4_41] = var_30_99;
var var_89_68 = "length";
var var_89_68 = window["history"]["length"];
var var_89_68 = XOR_STR(var_89_68, var_55_63);
var var_89_68 = btoa("" + var_89_68);
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "42.93", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_31_65 = 7.1;
var_82_53[var_31_65] = var_89_68;
var var_67_47 = 47.64;
var var_12_84 = [];
var var_71_99 = "__reactRouterContext";
var var_71_99 = window[var_71_99];
var var_20_46 = "state";
try { mem[var_20_46] = var_71_99[var_20_46]; } catch(r) { var_12_84 = "" + r; }
var var_76_26 = "loaderData";
var var_67_47 = XOR_STR(var_67_47, var_67_47);
var var_67_47 = btoa("" + var_67_47);
try { mem[var_76_26] = var_20_46[var_76_26]; } catch(r) { var_12_84 = "" + r; }
var var_57_61 = "root";
try { mem[var_57_61] = var_76_26[var_57_61]; } catch(r) { var_12_84 = "" + r; }
var var_96_5 = "clientBootstrap";
try { mem[var_96_5] = var_57_61[var_96_5]; } catch(r) { var_12_84 = "" + r; }
var var_88_19 = "cfConnectingIp";
try { mem[var_88_19] = var_96_5[var_88_19]; } catch(r) { var_12_84 = "" + r; }
var var_67_47 = XOR_STR(var_67_47, var_67_47);
var var_67_47 = btoa("" + var_67_47);
var var_12_84 = Array.isArray(var_12_84) ? (var_12_84.push(var_88_19), var_12_84) : var_12_84 + var_88_19;
var var_2_75 = "cfIpCity";
try { mem[var_2_75] = var_96_5[var_2_75]; } catch(r) { var_12_84 = "" + r; }
var var_21_21 = 97.61;
var var_67_47 = btoa("" + var_67_47);
var_82_53[var_21_21] = var_67_47;
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "81.04", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_12_84 = Array.isArray(var_12_84) ? (var_12_84.push(var_2_75), var_12_84) : var_12_84 + var_2_75;
var var_67_47 = 40.24;
var var_5_5 = "userRegion";
try { mem[var_5_5] = var_96_5[var_5_5]; } catch(r) { var_12_84 = "" + r; }
var var_12_84 = Array.isArray(var_12_84) ? (var_12_84.push(var_5_5), var_12_84) : var_12_84 + var_5_5;
var var_60_65 = "cfIpLatitude";
try { mem[var_60_65] = var_96_5[var_60_65]; } catch(r) { var_12_84 = "" + r; }
var var_12_84 = Array.isArray(var_12_84) ? (var_12_84.push(var_60_65), var_12_84) : var_12_84 + var_60_65;
var var_67_47 = 25.65;
var var_98_73 = "cfIpLongitude";
var var_21_21 = 97.61;
var var_67_47 = btoa("" + var_67_47);
var_82_53[var_21_21] = var_67_47;
try { mem[var_98_73] = var_96_5[var_98_73]; } catch(r) { var_12_84 = "" + r; }
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "37.84", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_12_84 = Array.isArray(var_12_84) ? (var_12_84.push(var_98_73), var_12_84) : var_12_84 + var_98_73;
var var_86_99 = 98.45;
var var_12_84 = JSON.stringify(var_12_84);
var var_12_84 = XOR_STR(var_12_84, var_55_63);
var var_12_84 = btoa("" + var_12_84);
var var_67_47 = 50.24;
var var_30_69 = 75.89;
var_82_53[var_30_69] = var_12_84;
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "36.65", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_20_41 = "location";
var var_20_41 = window["document"]["location"];
var var_19_37 = "";
var var_19_37 = Array.isArray(var_19_37) ? (var_19_37.push(var_20_41), var_19_37) : var_19_37 + var_20_41;
var var_94_65 = 36.78;
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "26.42", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_19_37 = XOR_STR(var_19_37, var_55_63);
var var_19_37 = btoa("" + var_19_37);
var var_67_72 = 84.91;
var_82_53[var_67_72] = var_19_37;
var var_2_63 = 30.7;
var var_10_03 = 27.36;
var var_50_45 = "random";
var var_50_45 = window["Math"]["random"];
var var_90_49 = var_50_45();
var var_90_49 = XOR_STR(var_90_49, var_90_49);
var var_86_99 = 94.12;
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "16.1", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_90_49 = btoa("" + var_90_49);
var_82_53[var_2_63] = var_90_49;
var var_87_58 = var_33_97();
Math.abs(var_43_93 - var_87_58) > var_61_37 ? var_7_75 = var_20_04 : null;
var var_7_75 = var_7_75 !== void 0 ? (mem["99.65"] = "33.47", var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (var_10_27(var_10_27, mem["82.53"], var_99_65, var_87_58) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (atob("" + var_7_75) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (XOR_STR(var_7_75, mem["42.48"]) || var_7_75) : var_7_75;
var var_7_75 = var_7_75 !== void 0 ? (JSON.parse(var_7_75) || var_7_75) : var_7_75;
var var_90_49 = var_50_45();
var_82_53[var_10_03] = var_90_49;
var var_44_95 = "setItem";
var var_44_95 = window["localStorage"]["setItem"].bind(window["localStorage"]);
var var_19_53 = "a757407bfe6f217d";
var var_94_65 = 9.28;
var var_3_39 = 4;
var var_66_06 = 73.37;
var var_94_65 = btoa("" + var_94_65);
var_82_53[var_66_06] = var_94_65;
var_44_95(var_19_53, var_3_39);
var var_20_04 = var_20_04 !== void 0 ? (mem["42.48"] = "29.18", var_20_04) : var_20_04;
var var_20_04 = var_20_04 !== void 0 ? (atob("" + var_20_04) || var_20_04) : var_20_04;
var var_20_04 = var_20_04 !== void 0 ? (XOR_STR(var_20_04, mem["42.48"]) || var_20_04) : var_20_04;
var var_20_04 = var_20_04 !== void 0 ? (JSON.parse(var_20_04) || var_20_04) : var_20_04;
var var_86_99 = 85.19;
var var_82_53 = JSON.stringify(var_82_53);
var var_94_65 = 90.33;
var var_82_53 = XOR_STR(var_82_53, var_55_63);
console.log(btoa("" + var_82_53));
</file>

<file path="manual.py">
from wrapper import ChatGPT

print(ChatGPT().ask_question("Test"))
</file>

<file path="readme.md">
# ChatGPT Reverse

A reverse-engineered implementation of ChatGPT that bypasses OpenAI's API system. This project provides free access to ChatGPT as API by emulating browser behavior and solving OpenAI Turnstile challenge through VM decompilation.

## Quick Start

**Install dependencies:**
```bash
pip install fastapi uvicorn curl-cffi pydantic pillow colorama esprima
```

**Run the API server:**
```bash
python api_server.py
```

The server runs on `http://localhost:6969`

**Or test directly:**
```bash
python manual.py
```

## Usage

### API Endpoint

POST `/conversation`

Request:
```json
{
    "proxy": "http://user:pass@ip:port",
    "message": "Your message here",
    "image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..."
}
```

The image field accepts base64-encoded images with data URI format. The system automatically handles image upload, processing, and attachment to the conversation.

Response:
```json
{
    "status": "success", 
    "result": "ChatGPT response"
}
```

### Direct Usage

```python
from wrapper import ChatGPT

client = ChatGPT("http://user:pass@proxy:port")
response = client.ask_question("Hello")
print(response)

# With image
response = client.ask_question("What's in this image?", "data:image/png;base64,...")
```
![alt text](images/image.png)

## Request Flow and Authentication

The system makes several authenticated requests to ChatGPT endpoints:

### 1. Initial Requirements (`/backend-anon/sentinel/chat-requirements`)
- Sends VM token in request body as `p` parameter
- Receives chat requirements token, proof-of-work challenge, and Turnstile bytecode
- Headers: `oai-client-version`, `oai-device-id`

### 2. Conduit Token (`/backend-anon/f/conversation/prepare`)
- Prepares conversation context
- Uses `x-conduit-token: no-token` header initially
- Returns conduit token for actual conversation

### 3. Main Conversation (`/backend-anon/f/conversation`)
- **Critical headers with VM components:**
  - `openai-sentinel-chat-requirements-token`: From requirements endpoint
  - `openai-sentinel-proof-token`: Solved proof-of-work challenge
  - `openai-sentinel-turnstile-token`: **VM-generated Turnstile token**
  - `x-conduit-token`: From conduit preparation
  - `oai-echo-logs`: Timing data for behavioral analysis


## VM and Turnstile Token System

The core of this implementation is the VM system that solves OpenAI Turnstile challenges used in the `openai-sentinel-turnstile-token` header.

### How It Works

ChatGPT uses Turnstile anti-bot protection that sends encrypted JavaScript bytecode to verify legitimate browsers. Our system:

1. **Receives bytecode** from ChatGPT servers in the requirements response
2. **Decompiles bytecode** into readable JavaScript operations  
3. **Parses decompiled output** to extract variable assignments and XOR keys
4. **Executes VM logic** by building browser fingerprint payload with XOR encryption

### Decompiler

The decompiler (`wrapper/reverse/decompiler.py`) translates Turnstile's custom bytecode:

```python
mapping = {
        "1": "XOR_STR",
        "2": "SET_VALUE",
        "3": "BTOA",
        "4": "BTOA_2",
        "5": "ADD_OR_PUSH",
        "6": "ARRAY_ACCESS",
        "7": "CALL",
        "8": "COPY",
        "10": "window",
        "11": "GET_SCRIPT_SRC",
        "12": "GET_MAP",
        "13": "TRY_CALL",
        "14": "JSON_PARSE",
        "15": "JSON_STRINGIFY",
        "17": "CALL_AND_SET",
        "18": "ATOB",
        "19": "BTOA_3",
        "20": "IF_EQUAL_CALL",
        "21": "IF_DIFF_CALL",
        "22": "TEMP_STACK_CALL",
        "23": "IF_DEFINED_CALL",
        "24": "BIND_METHOD",
        "27": "REMOVE_OR_SUBTRACT",
        "28": "undefined",
        "25": "undefined",
        "26": "undefined",
        "29": "LESS_THAN",
        "31": "INCREMENT",
        "32": "DECREMENT_AND_EXEC",
        "33": "MULTIPLY"
    }
```

It handles:
- Multi-layer bytecode decryption using XOR operations
- Variable state tracking across operations
- Browser API simulation (window, document objects)
- Three-stage decompilation process with nested encrypted layers

Also check decompiled.js for the full decompiled VM output.

![alt text](images/image-1.png)

### VM Execution

The VM (`wrapper/reverse/vm.py`) parses the decompiled output and extracts:

- **XOR encryption keys**: Dynamic keys for payload encryption extracted from `XOR_STR()` patterns
- **Variable assignments**: Browser fingerprint data collection points
- **Operation sequences**: DOM manipulation and measurement operations

Looking at the decompiled code, the VM identifies patterns like:
- `XOR_STR()` calls with different variables for encryption
- `getBoundingClientRect()` for element measurements  
- `navigator` property access for hardware fingerprints
- `localStorage.keys()` for browser storage data
- `Math.random()` calls requiring XOR self-encryption
- Performance timing measurements

The VM then builds the payload by providing realistic values:
- **Element measurements**: Realistic `getBoundingClientRect()` values
- **Navigator data**: Chrome browser fingerprints (`["Google Inc.","Win32",8,0]`)
- **Location info**: Current page URL and IP geolocation data
- **Storage keys**: Realistic localStorage key patterns
- **Random values**: Self-XOR encrypted random numbers

![alt text](images/image-2.png)

### Proof-of-Work Challenge

In addition to Turnstile, the system solves computational proof-of-work challenges:

```python
def solve_pow(seed, difficulty, config):
    for i in range(500000):
        result = runCheck(seed, difficulty, i, config)
        if result:
            return f"gAAAAAB{result}"
```

The PoW solver:
- Brute forces hash computations until difficulty threshold is met
- Uses FNV-1a hash algorithm with specific bit operations
- Integrates browser configuration data into the solution
- Provides the `openai-sentinel-proof-token` header value

### Challenge Integration

All challenge solutions work together:
1. **VM Token**: Generated from initial config, sent to `/chat-requirements`
2. **PoW Solution**: Computed from returned challenge, proves computational work
3. **Turnstile Token**: VM-generated from bytecode, proves browser legitimacy
4. **Conduit Token**: Session-specific token for conversation context

![alt text](images/image-3.png)

## Image Processing

The system handles image uploads through a multi-step process:

1. **Image Upload** (`/backend-anon/files`):
   - Accepts base64 image data
   - Returns file ID and Azure blob upload URL

2. **Blob Upload**: 
   - Uploads raw image data to Azure storage
   - Uses specific headers: `x-ms-blob-type: BlockBlob`

3. **Processing** (`/backend-anon/files/process_upload_stream`):
   - Processes uploaded image for multimodal use
   - Returns confirmation when ready

4. **Conversation Integration**:
   - References uploaded image via `file-service://{file_id}`
   - Includes image metadata (dimensions, mime type)

## Security Bypasses

**IP Detection**: Uses proxy rotation to avoid rate limits and geographic restrictions
**Browser Fingerprinting**: Emulates Chrome 139 with consistent headers and device properties
**Device Tracking**: Maintains stable device IDs (`oai-did`) across sessions  
**Behavioral Analysis**: Simulates realistic timing patterns in `oai-echo-logs`
**Challenge Solving**: Passes all cryptographic and computational verification steps

## Configuration

Proxy format: `http://username:password@host:port`

![alt text](images/image-4.png)

## Troubleshooting

**"Unusual activity" error**: Your IP is flagged, use a different proxy
**Requirements token fails**: VM token generation issue, check decompiler mappings  
**Turnstile token invalid**: Bytecode format changed, update operation mappings
**PoW challenge timeout**: Increase iteration limit or check hash implementation
**Image upload fails**: Check file size limits and blob storage connectivity

## Disclaimer

This project is for educational purposes only. If OpenAI has an Issue with this Project please contact me via my email nuhuh3116@gmail.com.
</file>

<file path="wrapper/__init__.py">
from .logger              import Log
from .runtime             import Run, Utils
from .IP_Info.headers     import Headers
from .reverse.challenges  import Challenges
from .reverse.parse       import Parser
from .reverse.vm          import VM
from .IP_Info.ip_info     import IP_Info
from .chatgpt             import ChatGPT
</file>

<file path="wrapper/chatgpt.py">
from wrapper      import Log, Utils, Headers, Challenges, VM, IP_Info
from random       import randint, random, choice
from zoneinfo     import ZoneInfo
from curl_cffi    import requests
from datetime     import datetime
from uuid         import uuid4
from json         import loads
from time         import time
from typing       import Any
from base64       import b64decode
from PIL import Image
from io import BytesIO


class ChatGPT:
    
    
    def __init__(self, proxy: str=None, cookies: dict = None) -> Any:
        self.session: requests.session.Session = requests.Session(impersonate="chrome133a")
        self.session.headers = Headers.DEFAULT
        self.data: dict = {}
        
        if proxy:
            
            self.session.proxies = {
                "all": proxy # format http://user:pass@ip:port
            }
            
        self.ip_info: list = IP_Info.fetch_info(self.session)
        self.timezone_offset: int = int(datetime.now(ZoneInfo(self.ip_info[5])).utcoffset().total_seconds() / 60)
        self.reacts: list = [
            "location",
            "__reactContainer$" + self._generate_react(),
            "_reactListening" + self._generate_react(),
        ]
        self.window_keys: list = [
            "0",
            "window",
            "self",
            "document",
            "name",
            "location",
            "customElements",
            "history",
            "navigation",
            "locationbar",
            "menubar",
            "personalbar",
            "scrollbars",
            "statusbar",
            "toolbar",
            "status",
            "closed",
            "frames",
            "length",
            "top",
            "opener",
            "parent",
            "frameElement",
            "navigator",
            "origin",
            "external",
            "screen",
            "innerWidth",
            "innerHeight",
            "scrollX",
            "pageXOffset",
            "scrollY",
            "pageYOffset",
            "visualViewport",
            "screenX",
            "screenY",
            "outerWidth",
            "outerHeight",
            "devicePixelRatio",
            "event",
            "clientInformation",
            "screenLeft",
            "screenTop",
            "styleMedia",
            "onsearch",
            "trustedTypes",
            "performance",
            "onappinstalled",
            "onbeforeinstallprompt",
            "crypto",
            "indexedDB",
            "sessionStorage",
            "localStorage",
            "onbeforexrselect",
            "onabort",
            "onbeforeinput",
            "onbeforematch",
            "onbeforetoggle",
            "onblur",
            "oncancel",
            "oncanplay",
            "oncanplaythrough",
            "onchange",
            "onclick",
            "onclose",
            "oncontentvisibilityautostatechange",
            "oncontextlost",
            "oncontextmenu",
            "oncontextrestored",
            "oncuechange",
            "ondblclick",
            "ondrag",
            "ondragend",
            "ondragenter",
            "ondragleave",
            "ondragover",
            "ondragstart",
            "ondrop",
            "ondurationchange",
            "onemptied",
            "onended",
            "onerror",
            "onfocus",
            "onformdata",
            "oninput",
            "oninvalid",
            "onkeydown",
            "onkeypress",
            "onkeyup",
            "onload",
            "onloadeddata",
            "onloadedmetadata",
            "onloadstart",
            "onmousedown",
            "onmouseenter",
            "onmouseleave",
            "onmousemove",
            "onmouseout",
            "onmouseover",
            "onmouseup",
            "onmousewheel",
            "onpause",
            "onplay",
            "onplaying",
            "onprogress",
            "onratechange",
            "onreset",
            "onresize",
            "onscroll",
            "onsecuritypolicyviolation",
            "onseeked",
            "onseeking",
            "onselect",
            "onslotchange",
            "onstalled",
            "onsubmit",
            "onsuspend",
            "ontimeupdate",
            "ontoggle",
            "onvolumechange",
            "onwaiting",
            "onwebkitanimationend",
            "onwebkitanimationiteration",
            "onwebkitanimationstart",
            "onwebkittransitionend",
            "onwheel",
            "onauxclick",
            "ongotpointercapture",
            "onlostpointercapture",
            "onpointerdown",
            "onpointermove",
            "onpointerrawupdate",
            "onpointerup",
            "onpointercancel",
            "onpointerover",
            "onpointerout",
            "onpointerenter",
            "onpointerleave",
            "onselectstart",
            "onselectionchange",
            "onanimationend",
            "onanimationiteration",
            "onanimationstart",
            "ontransitionrun",
            "ontransitionstart",
            "ontransitionend",
            "ontransitioncancel",
            "onafterprint",
            "onbeforeprint",
            "onbeforeunload",
            "onhashchange",
            "onlanguagechange",
            "onmessage",
            "onmessageerror",
            "onoffline",
            "ononline",
            "onpagehide",
            "onpageshow",
            "onpopstate",
            "onrejectionhandled",
            "onstorage",
            "onunhandledrejection",
            "onunload",
            "isSecureContext",
            "crossOriginIsolated",
            "scheduler",
            "alert",
            "atob",
            "blur",
            "btoa",
            "cancelAnimationFrame",
            "cancelIdleCallback",
            "captureEvents",
            "clearInterval",
            "clearTimeout",
            "close",
            "confirm",
            "createImageBitmap",
            "fetch",
            "find",
            "focus",
            "getComputedStyle",
            "getSelection",
            "matchMedia",
            "moveBy",
            "moveTo",
            "open",
            "postMessage",
            "print",
            "prompt",
            "queueMicrotask",
            "releaseEvents",
            "reportError",
            "requestAnimationFrame",
            "requestIdleCallback",
            "resizeBy",
            "resizeTo",
            "scroll",
            "scrollBy",
            "scrollTo",
            "setInterval",
            "setTimeout",
            "stop",
            "structuredClone",
            "webkitCancelAnimationFrame",
            "webkitRequestAnimationFrame",
            "chrome",
            "caches",
            "cookieStore",
            "ondevicemotion",
            "ondeviceorientation",
            "ondeviceorientationabsolute",
            "sharedStorage",
            "documentPictureInPicture",
            "fetchLater",
            "getScreenDetails",
            "queryLocalFonts",
            "showDirectoryPicker",
            "showOpenFilePicker",
            "showSaveFilePicker",
            "originAgentCluster",
            "viewport",
            "onpageswap",
            "onpagereveal",
            "credentialless",
            "fence",
            "launchQueue",
            "speechSynthesis",
            "oncommand",
            "onscrollend",
            "onscrollsnapchange",
            "onscrollsnapchanging",
            "webkitRequestFileSystem",
            "webkitResolveLocalFileSystemURL",
            "define",
            "ethereum",
            "__oai_SSR_HTML",
            "__reactRouterContext",
            "$RC",
            "__oai_SSR_TTI",
            "__reactRouterManifest",
            "__reactRouterVersion",
            "DD_RUM",
            "__REACT_INTL_CONTEXT__",
            "regeneratorRuntime",
            "DD_LOGS",
            "__STATSIG__",
            "__mobxInstanceCount",
            "__mobxGlobals",
            "_g",
            "__reactRouterRouteModules",
            "__SEGMENT_INSPECTOR__",
            "__reactRouterDataRouter",
            "MotionIsMounted",
            "_oaiHandleSessionExpired"
        ]
        
        if not cookies:
            self._fetch_cookies()
        else:
            self.session.cookies.update(cookies)
            
    def _generate_react(self) -> str:
        n = random() 
        base36 = ''
        chars = '0123456789abcdefghijklmnopqrstuvwxyz'
        x = int(n * 36**10)
        for _ in range(10):
            x, r = divmod(x, 36)
            base36 = chars[r] + base36
        return base36
    
    def _parse_event_stream(self, stream_data: str) -> str:
        result: list = []
        lines: list = stream_data.strip().split('\n')
        
        for line in lines:
            if line.startswith('data:'):
                
                data_str: str = line[5:].strip()
                
                if data_str == '[DONE]':
                    break
                
                data: dict = loads(data_str)
                
                if isinstance(data, dict):
                    
                    if data.get('o') == 'append' and data.get('p') == '/message/content/parts/0':
                        
                        result.append(data.get('v'))
                        
                    elif data.get('o') == 'patch' and isinstance(data.get('v'), list):
                        
                        for op in data.get('v'):
                            
                            if op.get('o') == 'append' and op.get('p') == '/message/content/parts/0':
                                
                                result.append(op.get('v'))
                                
                    elif 'v' in data and isinstance(data['v'], str):
                        result.append(data['v'])
                        
        return (''.join(result)).replace("\n", "")
        
    def _fetch_cookies(self) -> None:
        
        load_site: requests.models.Response = self.session.get("https://chatgpt.com")
        self.session.cookies.update(load_site.cookies)

        self.data["prod"] = load_site.text.split('data-build="')[1].split('"')[0]
        self.data["device-id"] = self.session.cookies.get("oai-did")
        
        self.start_time: int = int(time() * 1000)
        self.sid: str = str(uuid4())
        
        self.data["config"] = [
            4880,
            datetime.now(ZoneInfo(self.ip_info[5])).strftime(f"%a %b %d %Y %H:%M:%S GMT%z ({datetime.now(ZoneInfo(self.ip_info[5])).tzname()})"),
            4294705152,
            random(),
            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            None,
            self.data["prod"],
            "de-DE",
            "de-DE,de,en-US,en",
            random(),
            "webkitGetUserMedia−function webkitGetUserMedia() { [native code] }",
            choice(self.reacts),
            choice(self.window_keys),
            randint(800, 1400) + random(),
            self.sid,
            "",
            20,
            self.start_time
        ]
    
    def _get_tokens(self, process_time: int=randint(1400, 2000)) -> None:
        
        self.session.headers = Headers.REQUIREMENTS
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
        })
        
        p_value: str = Challenges.generate_token(self.data["config"])
        self.data["vm_token"] = p_value
        
        self.data["config"] = [
            4880,
            datetime.now(ZoneInfo(self.ip_info[5])).strftime(f"%a %b %d %Y %H:%M:%S GMT%z ({datetime.now(ZoneInfo(self.ip_info[5])).tzname()})"),
            4294705152,
            random(),
            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            None,
            self.data["prod"],
            "de-DE",
            "de-DE,de,en-US,en",
            random(),
            "webkitGetUserMedia−function webkitGetUserMedia() { [native code] }",
            choice(self.reacts),
            choice(self.window_keys),
            process_time + random(),
            self.sid,
            "",
            20,
            self.start_time
        ]
        
        requirements_data: dict = {
            'p': p_value,
        }
        
        requirements_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/sentinel/chat-requirements', json=requirements_data)

        if requirements_request.status_code == 200:
            self.data["token"] = requirements_request.json().get("token")
            self.data["proofofwork"] = requirements_request.json().get("proofofwork")
            self.data["bytecode"] = requirements_request.json().get("turnstile").get("dx")
        
        else:
            Log.Error("Something went wrong while fetching chat requirements")
    
    def get_conduit(self, next: bool = False) -> str:
        self.session.headers = Headers.CONDUIT
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
        })
        
        if not next:
            post_data: dict = {
                'action': 'next',
                'fork_from_shared_post': False,
                'parent_message_id': 'client-created-root',
                'model': 'auto',
                'timezone_offset_min': self.timezone_offset,
                'timezone': self.ip_info[5],
                'history_and_training_disabled': True,
                'conversation_mode': {
                    'kind': 'primary_assistant',
                },
                'system_hints': [],
                'supports_buffering': True,
                'supported_encodings': [
                    'v1',
                ],
            }
        
        else:
            post_data: dict = {
                'action': 'next',
                'fork_from_shared_post': False,
                'conversation_id': self.data["conversation_id"],
                'parent_message_id': self.data["parent_message_id"],
                'model': 'auto',
                'timezone_offset_min': self.timezone_offset,
                'timezone': self.ip_info[5],
                'history_and_training_disabled': True,
                'conversation_mode': {
                    'kind': 'primary_assistant',
                },
                'system_hints': [],
                'supports_buffering': True,
                'supported_encodings': [
                    'v1',
                ],
            }
                    
        conduit_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/f/conversation/prepare', json=post_data)
        
        if '"status":"ok"' in conduit_request.text:
            return conduit_request.json().get("conduit_token")
        
        else:
            Log.Error("Something went wrong while fetching conduit token: ")
            Log.Error(conduit_request.text)
            return None
    
    def start_conversation(self, message: str) -> None:
        
        self._get_tokens()
        conduit_token: str = self.get_conduit()
        
        time_1: int = randint(6000, 9000)
        proof_token: str = Challenges.solve_pow(self.data["proofofwork"]["seed"], self.data["proofofwork"]["difficulty"], self.data["config"])
        Log.Success(f"Solved POW: {proof_token}")
        turnstile_token: str = VM.get_turnstile(self.data["bytecode"], self.data["vm_token"], str(self.ip_info[:-1]))

        self.session.headers = Headers.CONVERSATION
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
            'oai-echo-logs': f'0,{time_1},1,{time_1 + randint(1000, 1200)}',
            'openai-sentinel-chat-requirements-token': self.data["token"],
            'openai-sentinel-proof-token': proof_token,
            'openai-sentinel-turnstile-token': turnstile_token,
            'x-conduit-token': conduit_token,
        })

        conversation_data: dict = {
            'action': 'next',
            'messages': [
                {
                    'id': str(uuid4()),
                    'author': {
                        'role': 'user',
                    },
                    'create_time': round(time(), 3),
                    'content': {
                        'content_type': 'text',
                        'parts': [
                            message,
                        ],
                    },
                    'metadata': {
                        'selected_github_repos': [],
                        'selected_all_github_repos': False,
                        'serialization_metadata': {
                            'custom_symbol_offsets': [],
                        },
                    },
                },
            ],
            'parent_message_id': 'client-created-root',
            'model': 'auto',
            'timezone_offset_min': self.timezone_offset,
            'timezone': self.ip_info[5],
            'history_and_training_disabled': True,
            'conversation_mode': {
                'kind': 'primary_assistant',
            },
            'enable_message_followups': True,
            'system_hints': [],
            'supports_buffering': True,
            'supported_encodings': [
                'v1',
            ],
            'client_contextual_info': {
                'is_dark_mode': True,
                'time_since_loaded': randint(3, 6),
                'page_height': 1219,
                'page_width': 3440,
                'pixel_ratio': 1,
                'screen_height': 1440,
                'screen_width': 3440,
            },
            'paragen_cot_summary_display_override': 'allow',
            'force_parallel_switch': 'auto',
        }
        
        conversation_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/f/conversation', json=conversation_data)
        self.session.cookies.update(conversation_request.cookies)
        
        if 'Unusual activity' in conversation_request.text:
            Log.Error("Your IP got flagged by chatgpt, retry with a new IP")
            exit(conversation_request.status_code)
        
        self.data["conversation_id"] = Utils.between(conversation_request.text, '"conversation_id": "', '"')
        self.data["parent_message_id"] = Utils.between(conversation_request.text, '"message_id": "', '"')
        self.response = self._parse_event_stream(conversation_request.text)

    def upload_image(self, image: str) -> None:
        
        self.session.headers = Headers.REQUIREMENTS
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
        })
        
        self.file_name: str = str(uuid4())
        
        if image.startswith("data:image"):
            image = image.split(",")[1]
            
        self.file_size: int = len(b64decode(image))
        self.width, self.height = Image.open(BytesIO(b64decode(image))).size
        
        image_data: dict = {
            'file_name': f'{self.file_name}.png',
            'file_size': self.file_size,
            'use_case': 'multimodal',
            'timezone_offset_min': self.timezone_offset,
            'reset_rate_limits': False,
        }
        file_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/files', json=image_data)
        
        self.data["file_id"] = file_request.json().get("file_id")
        upload_url: str = file_request.json().get("upload_url")
        
        self.session.headers = Headers.FILE
        upload_request: requests.models.Response = self.session.put(upload_url, data=b64decode(image))

        self.session.headers = Headers.REQUIREMENTS
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
        })
        
        process_data: dict = {
            'file_id': self.data["file_id"],
            'use_case': 'multimodal',
            'index_for_retrieval': False,
            'file_name': f'{self.file_name}.png',
        }
        
        process_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/files/process_upload_stream', json=process_data)
        
        if "Succeeded processing " in process_request.text:
            return
        else:
            Log.Error("Something went wrong while uploading image")
        
        
        
    def start_with_image(self, message: str, image: str) -> None:
        
        self._get_tokens()
        conduit_token: str = self.get_conduit()
        self.upload_image(image)
        
        time_1: int = randint(6000, 9000)
        proof_token: str = Challenges.solve_pow(self.data["proofofwork"]["seed"], self.data["proofofwork"]["difficulty"], self.data["config"])
        
        turnstile_token: str = VM.get_turnstile(self.data["bytecode"], self.data["vm_token"], str(self.ip_info[:-1]))

        self.session.headers = Headers.CONVERSATION
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
            'oai-echo-logs': f'0,{time_1},1,{time_1 + randint(1000, 1200)}',
            'openai-sentinel-chat-requirements-token': self.data["token"],
            'openai-sentinel-proof-token': proof_token,
            'openai-sentinel-turnstile-token': turnstile_token,
            'x-conduit-token': conduit_token,
        })

        conversation_data: dict = {
            'action': 'next',
            'messages': [
                {
                    'id': str(uuid4()),
                    'author': {
                        'role': 'user',
                    },
                    'create_time': round(time(), 3),
                    'content': {
                        'content_type': 'multimodal_text',
                        'parts': [
                            {
                                'content_type': 'image_asset_pointer',
                                'asset_pointer': f'file-service://{self.data["file_id"]}',
                                'size_bytes': self.file_size,
                                'width': self.width,
                                'height': self.height,
                            },
                            message,
                        ],
                    },
                    'metadata': {
                        'attachments': [
                            {
                                'id': self.data["file_id"],
                                'size': self.file_size,
                                'name': f'{self.file_name}.png',
                                'mime_type': 'image/png',
                                'width': self.width,
                                'height': self.height,
                                'source': 'local',
                            },
                        ],
                        'selected_github_repos': [],
                        'selected_all_github_repos': False,
                        'serialization_metadata': {
                            'custom_symbol_offsets': [],
                        },
                    },
                },
            ],
            'parent_message_id': 'client-created-root',
            'model': 'auto',
            'timezone_offset_min': self.timezone_offset,
            'timezone': self.ip_info[5],
            'history_and_training_disabled': True,
            'conversation_mode': {
                'kind': 'primary_assistant',
            },
            'enable_message_followups': True,
            'system_hints': [],
            'supports_buffering': True,
            'supported_encodings': [
                'v1',
            ],
            'client_contextual_info': {
                'is_dark_mode': True,
                'time_since_loaded': randint(3, 6),
                'page_height': 1219,
                'page_width': 3440,
                'pixel_ratio': 1,
                'screen_height': 1440,
                'screen_width': 3440,
            },
            'paragen_cot_summary_display_override': 'allow',
            'force_parallel_switch': 'auto',
        }
        
        conversation_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/f/conversation', json=conversation_data)
        self.session.cookies.update(conversation_request.cookies)
        
        if 'Unusual activity' in conversation_request.text:
            Log.Error("Your IP got flagged by chatgpt, retry with a new IP")
            exit(conversation_request.status_code)
        
        self.data["conversation_id"] = Utils.between(conversation_request.text, '"conversation_id": "', '"')
        self.data["parent_message_id"] = Utils.between(conversation_request.text, '"message_id": "', '"')
        self.response = self._parse_event_stream(conversation_request.text)
    
    def hold_conversation(self, message: str, new: bool = True) -> None:
        self.index = 2000
        
        if new:
            self.start_conversation(message)
        
        conduit_token: str = self.get_conduit(next=True)
        
        self._get_tokens(randint(self.index, self.index + 1000))
        self.index += 3000
        
        time_1: int = randint(self.index, self.index + 3000)
        proof_token: str = Challenges.solve_pow(self.data["proofofwork"]["seed"], self.data["proofofwork"]["difficulty"], self.data["config"])
        
        turnstile_token: str = VM.get_turnstile(self.data["bytecode"], self.data["vm_token"], str(self.ip_info[:-1]))


        self.session.headers = Headers.CONVERSATION
        self.session.headers.update({
            'oai-client-version': self.data["prod"],
            'oai-device-id': self.data["device-id"],
            'oai-echo-logs': f'0,{time_1},1,{time_1 + randint(1000, 1200)}',
            'openai-sentinel-chat-requirements-token': self.data["token"],
            'openai-sentinel-proof-token': proof_token,
            'openai-sentinel-turnstile-token': turnstile_token,
            'x-conduit-token': conduit_token,
        })
        
        if new:
            new_message: str = input("Prompt: ")
        else:
            new_message: str = message
        
        conversation_data: dict = {
            'action': 'next',
            'messages': [
                {
                    'id': str(uuid4()),
                    'author': {
                        'role': 'user',
                    },
                    'create_time': round(time(), 3),
                    'content': {
                        'content_type': 'text',
                        'parts': [
                            new_message,
                        ],
                    },
                    'metadata': {
                        'selected_github_repos': [],
                        'selected_all_github_repos': False,
                        'serialization_metadata': {
                            'custom_symbol_offsets': [],
                        },
                    },
                },
            ],
            'conversation_id': self.data["conversation_id"],
            'parent_message_id': self.data["parent_message_id"],
            'model': 'auto',
            'timezone_offset_min': self.timezone_offset,
            'timezone': self.ip_info[5],
            'history_and_training_disabled': True,
            'conversation_mode': {
                'kind': 'primary_assistant',
            },
            'enable_message_followups': True,
            'system_hints': [],
            'supports_buffering': True,
            'supported_encodings': [
                'v1',
            ],
            'client_contextual_info': {
                'is_dark_mode': True,
                'time_since_loaded': 17,
                'page_height': 1219,
                'page_width': 3440,
                'pixel_ratio': 1,
                'screen_height': 1440,
                'screen_width': 3440,
            },
            'paragen_cot_summary_display_override': 'allow',
            'force_parallel_switch': 'auto',
        }
        
        conversation_request: requests.models.Response = self.session.post('https://chatgpt.com/backend-anon/f/conversation', json=conversation_data)
        self.session.cookies.update(conversation_request.cookies)
        
        if 'Unusual activity' in conversation_request.text:
            Log.Error("Your IP got flagged by chatgpt, retry with a new IP")
            exit(conversation_request.status_code)
        
        self.data["conversation_id"] = Utils.between(conversation_request.text, '"conversation_id": "', '"')
        self.data["parent_message_id"] = Utils.between(conversation_request.text, '"message_id": "', '"')
        
        self.response = self._parse_event_stream(conversation_request.text)
    
    def ask_question(self, message: str, image: str = None) -> str:
        
        if not image:
            self.start_conversation(message)
        else:
            self.start_with_image(message, image)
        
        return self.response
</file>

<file path="wrapper/IP_Info/headers.py">
class Headers:
    
    DEFAULT: dict = {
        'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7',
        'accept-language': 'de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7',
        'cache-control': 'no-cache',
        'pragma': 'no-cache',
        'priority': 'u=0, i',
        'sec-ch-ua': '"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"Windows"',
        'sec-fetch-dest': 'document',
        'sec-fetch-mode': 'navigate',
        'sec-fetch-site': 'same-origin',
        'sec-fetch-user': '?1',
        'upgrade-insecure-requests': '1',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36',
    }
    
    REQUIREMENTS: dict = {
        'accept': '*/*',
        'accept-language': 'de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7',
        'cache-control': 'no-cache',
        'content-type': 'application/json',
        'oai-client-version': '',
        'oai-device-id': '',
        'oai-language': 'de-DE',
        'origin': 'https://chatgpt.com',
        'pragma': 'no-cache',
        'priority': 'u=1, i',
        'referer': 'https://chatgpt.com/',
        'sec-ch-ua': '"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"Windows"',
        'sec-fetch-dest': 'empty',
        'sec-fetch-mode': 'cors',
        'sec-fetch-site': 'same-origin',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36',
    }
    
    CONDUIT: dict = {
        'accept': '*/*',
        'accept-language': 'de-DE,de;q=0.9',
        'cache-control': 'no-cache',
        'content-type': 'application/json',
        'oai-client-version': '',
        'oai-device-id': '',
        'oai-language': 'de-DE',
        'origin': 'https://chatgpt.com',
        'pragma': 'no-cache',
        'priority': 'u=1, i',
        'referer': 'https://chatgpt.com/',
        'sec-ch-ua': '"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"Windows"',
        'sec-fetch-dest': 'empty',
        'sec-fetch-mode': 'cors',
        'sec-fetch-site': 'same-origin',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36',
        'x-conduit-token': 'no-token',
    }
    
    CONVERSATION: dict = {
        'accept': 'text/event-stream',
        'accept-language': 'de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7',
        'cache-control': 'no-cache',
        'content-type': 'application/json',
        'oai-client-version': '',
        'oai-device-id': '',
        'oai-echo-logs': '',
        'oai-language': 'de-DE',
        'openai-sentinel-chat-requirements-token': '',
        'openai-sentinel-proof-token': '',
        'openai-sentinel-turnstile-token': '',
        'origin': 'https://chatgpt.com',
        'pragma': 'no-cache',
        'priority': 'u=1, i',
        'referer': 'https://chatgpt.com/',
        'sec-ch-ua': '"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"Windows"',
        'sec-fetch-dest': 'empty',
        'sec-fetch-mode': 'cors',
        'sec-fetch-site': 'same-origin',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36',
        'x-conduit-token': '',
    }
    
    FILE: dict = {
        'accept': 'application/json, text/plain, */*',
        'accept-language': 'de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7',
        'cache-control': 'no-cache',
        'content-type': 'image/png',
        'origin': 'https://chatgpt.com',
        'pragma': 'no-cache',
        'priority': 'u=1, i',
        'referer': 'https://chatgpt.com/',
        'sec-ch-ua': '"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"Windows"',
        'sec-fetch-dest': 'empty',
        'sec-fetch-mode': 'cors',
        'sec-fetch-site': 'cross-site',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36',
        'x-ms-blob-type': 'BlockBlob',
        'x-ms-version': '2020-04-08',
    }
</file>

<file path="wrapper/IP_Info/ip_info.py">
from curl_cffi  import requests
from ..logger   import Log
from ..runtime  import Utils


class IP_Info:
    
    @staticmethod
    def fetch_info(session: requests.session.Session) -> str:
        
        ip_infos: list = []
        
        info_request: requests.models.Response = session.get('https://iplocation.com/')
        
        ip_infos.append(Utils.between(info_request.text, '<td><b class="ip">', '<'))
        ip_infos.append(Utils.between(info_request.text, '<td class="city">', '<'))
        ip_infos.append(Utils.between(info_request.text, '<td><span class="region_name">', '<'))
        ip_infos.append(Utils.between(info_request.text, '<td class="lat">', '<'))
        ip_infos.append(Utils.between(info_request.text, '<td class="lng">', '<'))
        
        info_request_2: requests.models.Response = session.get('https://ipaddresslocation.net/ip-to-timezone')
        
        ip_infos.append(Utils.between(info_request_2.text, 'Time Zone:</strong> ', ' '))

        return ip_infos
</file>

<file path="wrapper/logger.py">
from typing      import Optional
from datetime    import datetime
from colorama    import Fore
from threading   import Lock
from time        import time


class Log:
    """
    Logging class to log text better in console.
    """
    
    colours: Optional[dict] = {
        'SUCCESS': Fore.LIGHTGREEN_EX,
        'ERROR': Fore.LIGHTRED_EX,
        'INFO': Fore.LIGHTWHITE_EX
    }
    
    lock = Lock()
    
    @staticmethod
    def _log(level, prefix, message) -> Optional[None]:
        """
        Private log function to build the payload to print.
        
        :param level: Just not used, only a filler
        :param prefix: Prefix to indicate if its Success, Error or Info
        :param message: Message to Log
        """
        
        timestamp: Optional[int] = datetime.fromtimestamp(time()).strftime("%H:%M:%S")
        
        log_message = (
            f"{Fore.LIGHTBLACK_EX}[{Fore.MAGENTA}{timestamp}{Fore.RESET}{Fore.LIGHTBLACK_EX}]{Fore.RESET} "
            f"{prefix} {message}"
        )
        
        with Log.lock:
            print(log_message)

    @staticmethod
    def Success(message, prefix="[+]", color=colours['SUCCESS']) -> Optional[None]:
        """
        Logging a Success message.
        """
        Log._log("SUCCESS", f"{color}{prefix}{Fore.RESET}", message)

    @staticmethod
    def Error(message, prefix="[!]", color=colours['ERROR']) -> Optional[None]:
        """
        Logging an Error Message.
        """
        Log._log("ERROR", f"{color}{prefix}{Fore.RESET}", message)

    @staticmethod
    def Info(message, prefix="[!]", color=colours['INFO']) -> Optional[None]:
        """
        Logging an Info Message.
        """
        Log._log("INFO", f"{color}{prefix}{Fore.RESET}", message)
</file>

<file path="wrapper/reverse/challenges.py">
from json   import dumps
from base64 import b64encode
from time   import time


class Challenges:


    @staticmethod
    def encode(e):
        e = dumps(e, separators=(",", ":")) 
        encoded = e.encode("utf-8")
        return b64encode(encoded).decode()

    @staticmethod
    def generate_token(config):
        t = "e"
        n = time() * 1000
        try:
            config[3] = 1
            config[9] = round(time() * 1000 - n)
            return "gAAAAAC" + Challenges.encode(config)
        except Exception as e:
            t = Challenges.encode(str(e))
        return "error_" + t
    
    @staticmethod
    def mod(e: str) -> str:
        t = 2166136261
        for ch in e:
            t ^= ord(ch)
            t = (t * 16777619) & 0xFFFFFFFF

        t ^= (t >> 16)
        t = (t * 2246822507) & 0xFFFFFFFF
        t ^= (t >> 13)
        t = (t * 3266489909) & 0xFFFFFFFF
        t ^= (t >> 16)

        return f"{t:08x}"

    @staticmethod
    def _runCheck(t0, n, r, o, config):
        config[3] = o
        config[9] = round(time() * 1000 - t0)

        i = Challenges.encode(config)

        if Challenges.mod(n + i)[:len(r)] <= r:
            return f"{i}~S"
        return None

    @staticmethod
    def solve_pow(t, n, config):
        t0 = int(time() * 1000)
        for i in range(500000):
            a = Challenges._runCheck(t0, t, n, i, config)
            if a:
                return "gAAAAAB" + a
</file>

<file path="wrapper/reverse/decompiler.py">
import re
import json
import base64


class Decompiler:
    
    
    mapping: dict = {
        "1": "XOR_STR",
        "2": "SET_VALUE",
        "3": "BTOA",
        "4": "BTOA_2",
        "5": "ADD_OR_PUSH",
        "6": "ARRAY_ACCESS",
        "7": "CALL",
        "8": "COPY",
        "10": "window",
        "11": "GET_SCRIPT_SRC",
        "12": "GET_MAP",
        "13": "TRY_CALL",
        "14": "JSON_PARSE",
        "15": "JSON_STRINGIFY",
        "17": "CALL_AND_SET",
        "18": "ATOB",
        "19": "BTOA_3",
        "20": "IF_EQUAL_CALL",
        "21": "IF_DIFF_CALL",
        "22": "TEMP_STACK_CALL",
        "23": "IF_DEFINED_CALL",
        "24": "BIND_METHOD",
        "27": "REMOVE_OR_SUBTRACT",
        "28": "undefined",
        "25": "undefined",
        "26": "undefined",
        "29": "LESS_THAN",
        "31": "INCREMENT",
        "32": "DECREMENT_AND_EXEC",
        "33": "MULTIPLY",
        "34": "MOVE"
    }

    functions: dict = {
            "XOR_STR": """function XOR_STR(e, t) {
        e = String(e);
        t = String(t);
        let n = "";
        for (let r = 0; r < e.length; r++)
            n += String.fromCharCode(e.charCodeAt(r) ^ t.charCodeAt(r % t.length));
        return n;
    }
    """
    }

    @staticmethod
    def start():
        Decompiler.xorkey = ""
        Decompiler.xorkey2 = ""
        Decompiler.decompiled = "var mem = {};\n"
        Decompiler.array_dict = {}
        Decompiler.vg = 0
        Decompiler.round1 = 0
        Decompiler.found = False
        Decompiler.potential = []

    @staticmethod
    def xS(e, t):
        n = ""
        for r in range(len(e)):
            n += chr(ord(e[r]) ^ ord(t[r % len(t)]))
        return n

    @staticmethod
    def handle_operation(operation, args):
        if operation == "COPY":
            Decompiler.mapping[args[0]] = Decompiler.mapping[args[1]]
            if Decompiler.mapping[args[1]] != "window":
                if Decompiler.mapping[args[1]] in Decompiler.functions and f"function {Decompiler.mapping[args[1]]}" not in Decompiler.decompiled:
                    Decompiler.decompiled += Decompiler.functions[Decompiler.mapping[args[1]]] + "\n"
            else:
                var_name = str(args[1]).replace(".", "_")
                Decompiler.decompiled += f"var var_{var_name} = window;\n"
                Decompiler.array_dict[args[1]] = "window"
        
        elif operation == "SET_VALUE":
            var_name = str(args[0]).replace(".", "_")
            value = args[1]
            try:
                num = float(value)
                if num.is_integer():
                    Decompiler.decompiled += f"var var_{var_name} = {int(num)};\n"
                    Decompiler.array_dict[args[0]] = str(int(num))
                else:
                    Decompiler.decompiled += f"var var_{var_name} = {num};\n"
                    Decompiler.array_dict[args[0]] = str(num)
            except (ValueError, TypeError):
                if isinstance(value, str):
                    if value == "[]":
                        Decompiler.decompiled += f"var var_{var_name} = [];\n"
                        Decompiler.array_dict[args[0]] = []
                    
                    elif value == "None":
                        Decompiler.decompiled += f"var var_{var_name} = null;\n"
                        Decompiler.array_dict[args[0]] = "null"
                        
                    else:
                        Decompiler.decompiled += f"var var_{var_name} = \"{value}\";\n"
                        Decompiler.array_dict[args[0]] = f"\"{value}\""
                elif isinstance(value, list):
                    Decompiler.decompiled += f"var var_{var_name} = [];\n"
                    Decompiler.array_dict[args[0]] = []
                elif value is None:
                    Decompiler.decompiled += f"var var_{var_name} = null;\n"
                    Decompiler.array_dict[args[0]] = "null"
                else:
                    Decompiler.decompiled += f"var var_{var_name} = {value};\n"
                    Decompiler.array_dict[args[0]] = str(value)
        
        elif operation == "ARRAY_ACCESS":
            Decompiler.handle_array_access(args)
        
        elif operation == "BIND_METHOD":
            Decompiler.handle_bind_method(args)
        
        elif operation == "XOR_STR":
            if Decompiler.round1 == 1 and len(Decompiler.potential) < 2:
                Decompiler.potential.append({"var": args[0], "key": args[1]})
            var_name = str(args[0]).replace(".", "_")
            key_name = str(args[1]).replace(".", "_")
            Decompiler.decompiled += f"var var_{var_name} = XOR_STR(var_{var_name}, var_{key_name});\n"
        
        elif operation == "BTOA_3":
            var_name = str(args[0]).replace(".", "_")
            Decompiler.decompiled += f"var var_{var_name} = btoa(\"\" + var_{var_name});\n"
        
        elif operation == "CALL_AND_SET":
            var_name = str(args[0]).replace(".", "_")
            func_name = str(args[1]).replace(".", "_")
            args_str = ", ".join(f"var_{arg.replace('.', '_')}" for arg in args[2:])
            Decompiler.decompiled += f"var var_{var_name} = var_{func_name}({args_str});\n"
        
        elif operation == "IF_DEFINED_CALL":
            Decompiler.handle_if_defined_call(args)
        
        elif operation == "CALL":
            Decompiler.handle_call_operation(args)
        
        elif operation == "ADD_OR_PUSH":
            var_name = str(args[0]).replace(".", "_")
            arg_name = str(args[1]).replace(".", "_")
            Decompiler.decompiled += (
                f"var var_{var_name} = Array.isArray(var_{var_name}) ? "
                f"(var_{var_name}.push(var_{arg_name}), var_{var_name}) : var_{var_name} + var_{arg_name};\n"
            )
        
        elif operation == "IF_DIFF_CALL":
            var_0 = str(args[0]).replace(".", "_")
            var_1 = str(args[1]).replace(".", "_")
            var_2 = str(args[2]).replace(".", "_")
            if Decompiler.mapping.get(args[3]) == "COPY":
                var_4 = str(args[4]).replace(".", "_")
                var_5 = str(args[5]).replace(".", "_")
                Decompiler.decompiled += (
                    f"Math.abs(var_{var_0} - var_{var_1}) > var_{var_2} ? var_{var_4} = var_{var_5} : null;\n"
                )
            else:
                args_str = ", ".join(f"var_{arg.replace('.', '_')}" for arg in args[4:])
                Decompiler.decompiled += (
                    f"Math.abs(var_{var_0} - var_{var_1}) > var_{var_2} ? {Decompiler.mapping[args[3]]}({args_str}) : null;\n"
                )
        
        elif operation == "TRY_CALL":
            Decompiler.handle_try_call(args)
        
        elif operation == "JSON_STRINGIFY":
            var_name = str(args[0]).replace(".", "_")
            Decompiler.decompiled += f"var var_{var_name} = JSON.stringify(var_{var_name});\n"
            
        elif operation == "MOVE":
            Decompiler.decompiled += f"MOVE {args}" # not even used lmfao
        
        else:
            mapped = [Decompiler.mapping[key] for key in args[1:] if key in Decompiler.mapping]
            unlabeled = [str(key) for key in args[1:] if key not in Decompiler.mapping]
            all_values = " ".join(mapped + unlabeled)
            Decompiler.decompiled += f"// UNKNOWN: {operation} -> {args[0]} {all_values};\n"

    @staticmethod
    def handle_try_call(args):
        target_var = f"var_{str(args[0]).replace('.', '_')}"
        fn = Decompiler.mapping.get(args[1], "")
        rest_args = [f"var_{str(a).replace('.', '_')}" for a in args[2:]]
        if fn == "ARRAY_ACCESS":
            Decompiler.decompiled += (
                f"try {{ mem[{rest_args[0]}] = {rest_args[1]}[{rest_args[0]}]; }} catch(r) {{ {target_var} = \"\" + r; }}\n"
            )
        else:
            args_str = ", ".join(rest_args)
            Decompiler.decompiled += (
                f"try {{ {fn}({args_str}); }} catch(r) {{ {target_var} = \"\" + r; }}\n"
            )

    @staticmethod
    def handle_array_access(args):
        var_0 = str(args[0]).replace(".", "_")
        var_1 = str(args[1]).replace(".", "_")
        var_2 = str(args[2]).replace(".", "_")
        if f"var var_{var_1} =" in Decompiler.decompiled:
            if args[1] in Decompiler.array_dict or args[2] in Decompiler.array_dict:
                if args[2] in Decompiler.array_dict and args[1] not in Decompiler.array_dict:
                    Decompiler.decompiled += f"var var_{var_0} = var_{var_1}[{Decompiler.array_dict[args[2]]}];\n"
                elif args[1] in Decompiler.array_dict and args[2] not in Decompiler.array_dict:
                    Decompiler.decompiled += f"var var_{var_0} = {Decompiler.array_dict[args[1]]}[var_{var_2}];\n"
                else:
                    if re.search(rf"var\s+var_{var_1}\s*=\s*\w+\([^)]*\)", Decompiler.decompiled):
                        Decompiler.decompiled += f"var var_{var_0} = var_{var_1}[{Decompiler.array_dict[args[2]]}];\n"
                        Decompiler.array_dict[args[0]] = f"var_{var_1}[{Decompiler.array_dict[args[2]]}]"
                    else:
                        Decompiler.decompiled += f"var var_{var_0} = {Decompiler.array_dict[args[1]]}[{Decompiler.array_dict[args[2]]}];\n"
                        Decompiler.array_dict[args[0]] = f"{Decompiler.array_dict[args[1]]}[{Decompiler.array_dict[args[2]]}]"
            else:
                Decompiler.decompiled += f"var var_{var_0} = var_{var_1}[var_{var_2}];\n"
        else:
            Decompiler.decompiled += f"var var_{var_0} = window[var_{var_2}];\n"

    @staticmethod
    def handle_bind_method(args):
        var_0 = str(args[0]).replace(".", "_")
        var_1 = str(args[1]).replace(".", "_")
        var_2 = str(args[2]).replace(".", "_")
        if f"var var_{var_1} =" in Decompiler.decompiled:
            if args[1] in Decompiler.array_dict or args[2] in Decompiler.array_dict:
                if args[1] in Decompiler.array_dict and args[2] not in Decompiler.array_dict:
                    Decompiler.decompiled += (
                        f"var var_{var_0} = {Decompiler.array_dict[args[1]]}[var_{var_2}].bind({Decompiler.array_dict[args[1]]});\n"
                    )
                else:
                    if re.search(rf"var\s+var_{var_1}\s*=\s*\w+\([^)]*\)", Decompiler.decompiled):
                        Decompiler.decompiled += (
                            f"var var_{var_0} = var_{var_1}[{Decompiler.array_dict[args[2]]}].bind(var_{var_1});\n"
                        )
                        Decompiler.array_dict[args[0]] = f"var_{var_1}[{Decompiler.array_dict[args[2]]}]"
                    else:
                        Decompiler.decompiled += (
                            f"var var_{var_0} = {Decompiler.array_dict[args[1]]}[{Decompiler.array_dict[args[2]]}].bind({Decompiler.array_dict[args[1]]});\n"
                        )
                        Decompiler.array_dict[args[0]] = f"{Decompiler.array_dict[args[1]]}[{Decompiler.array_dict[args[2]]}]"
            else:
                Decompiler.decompiled += (
                    f"var var_{var_0} = var_{var_1}[var_{var_2}].bind(var_{var_1});\n"
                )
        else:
            Decompiler.decompiled += (
                f"var var_{var_0} = window[var_{var_2}].bind(var_{var_1});\n"
            )

    @staticmethod
    def handle_if_defined_call(args):
        result = []
        for item in args:
            if item in Decompiler.mapping:
                keys = [k for k, v in Decompiler.mapping.items() if v == Decompiler.mapping[item] and k != item]
                result.append(keys[0] if keys else None)
            else:
                result.append(None)
        result = [
            None if key is None else ([k for k, v in Decompiler.mapping.items() if v == Decompiler.mapping[key] and k != key] or [None])[0]
            for key in result
        ]
        
        if len(args) == 4:
            target = str(args[3]).replace(".", "_")
            count = len(re.findall(target, Decompiler.decompiled))
            if count <= 1 and f"var var_{str(args[2]).replace('.', '_')}" not in Decompiler.decompiled:
                if not Decompiler.xorkey:
                    Decompiler.xorkey = str(args[3])
                var_0 = str(args[0]).replace(".", "_")
                arg_2 = str(args[2]).replace(".", "_")
                arg_3 = str(args[3]).replace(".", "_")
                if Decompiler.mapping.get(result[1]) == "SET_VALUE":
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? (mem[\"{args[2]}\"] = \"{args[3]}\", var_{var_0}) : var_{var_0};\n"
                    )
                else:
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? ({Decompiler.mapping[result[1]]}(\"{args[2]}\", \"{args[3]}\") || var_{var_0}) : var_{var_0};\n"
                    )
            elif count <= 3:
                var_0 = str(args[0]).replace(".", "_")
                arg_2 = str(args[2]).replace(".", "_")
                arg_3 = str(args[3]).replace(".", "_")
                if Decompiler.mapping.get(result[1]) == "SET_VALUE":
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? ((mem[\"{args[2]}\"] = \"{args[3]}\") || var_{var_0}) : var_{var_0};\n"
                    )
                else:
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? ({Decompiler.mapping[result[1]]}(var_{arg_2}, mem[\"{args[3]}\"]) || var_{var_0}) : var_{var_0};\n"
                    )
            elif Decompiler.mapping.get(result[1]) == "JSON_PARSE":
                var_0 = str(args[0]).replace(".", "_")
                arg_3 = str(args[3]).replace(".", "_")
                Decompiler.decompiled += (
                    f"var var_{var_0} = var_{var_0} !== void 0 ? (JSON.parse(var_{arg_3}) || var_{var_0}) : var_{var_0};\n"
                )
            else:
                var_0 = str(args[0]).replace(".", "_")
                args_str = ", ".join(f"var_{arg.replace('.', '_')}" for arg in args[2:])
                Decompiler.decompiled += (
                    f"var var_{var_0} = var_{var_0} !== void 0 ? ({Decompiler.mapping[result[1]]}({args_str}) || var_{var_0}) : var_{var_0};\n"
                )
        else:
            var_0 = str(args[0]).replace(".", "_")
            if len(args) > 4 and f"mem[\"{args[4]}\"] =" in Decompiler.decompiled:
                args_str = ", ".join(
                    f"mem[\"{arg}\"]" if i + 2 == 3 else f"var_{str(arg).replace('.', '_')}"
                    for i, arg in enumerate(args[2:])
                )
                if Decompiler.mapping.get(result[1]) == "CALL":
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? (var_{str(args[2]).replace('.', '_')}({args_str}) || var_{var_0}) : var_{var_0};\n"
                    )
                else:
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? ({Decompiler.mapping[result[1]]}({args_str}) || var_{var_0}) : var_{var_0};\n"
                    )
            else:
                args_str = ", ".join(f"var_{arg.replace('.', '_')}" for arg in args[2:])
                if Decompiler.mapping.get(result[1]) == "ATOB":
                    arg_2 = str(args[2]).replace(".", "_")
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? (atob(\"\" + var_{arg_2}) || var_{var_0}) : var_{var_0};\n"
                    )
                elif len(args) >= 3 and result[1] in Decompiler.mapping:
                    Decompiler.decompiled += (
                        f"var var_{var_0} = var_{var_0} !== void 0 ? ({Decompiler.mapping[result[1]]}({args_str}) || var_{var_0}) : var_{var_0};\n"
                    )
                else:
                    Decompiler.decompiled += f"// ERROR: Invalid IF_DEFINED_CALL with args {args};\n"

    @staticmethod
    def handle_call_operation(args):
        if args[0] in Decompiler.mapping:
            if Decompiler.mapping[args[0]] == "BTOA":
                arg_1 = str(args[1]).replace(".", "_")
                Decompiler.decompiled += f"console.log(btoa(\"\" + var_{arg_1}));\n"
            else:
                args_str = ", ".join(f"var_{arg.replace('.', '_')}" for arg in args)
                Decompiler.decompiled += f"{Decompiler.mapping[args[0]]}({args_str});\n"
        else:
            if f"var var_{str(args[0]).replace('.', '_')} = \"set\";" in Decompiler.decompiled:
                arg_1 = str(args[1]).replace(".", "_")
                arg_2 = str(args[2]).replace(".", "_")
                arg_3 = str(args[3]).replace(".", "_")
                Decompiler.decompiled += f"var_{arg_1}[var_{arg_2}] = var_{arg_3};\n"
            else:
                args_str = ", ".join(f"var_{arg.replace('.', '_')}" for arg in args[1:])
                Decompiler.decompiled += f"var_{str(args[0]).replace('.', '_')}({args_str});\n"

    @staticmethod
    def remove_unused_variables():
        lines = Decompiler.decompiled.split("\n")
        used_vars = set()
        var_decl_lines = []

        for i, line in enumerate(lines):
            match = re.match(r"^var\s+var_([\w_]+)\s*=", line)
            if match:
                var_decl_lines.append({"name": match.group(1), "index": i})

        for var in var_decl_lines:
            name = var["name"]
            is_used = any(name in line and not line.startswith(f"var var_{name} =") for line in lines)
            if is_used:
                used_vars.add(name)

        Decompiler.decompiled = "\n".join(
            line for line in lines
            if not re.match(r"^var\s+var_([\w_]+)\s*=", line) or re.match(r"^var\s+var_([\w_]+)\s*=", line).group(1) in used_vars
        )

    @staticmethod
    def decompile(bytecode):
        while len(bytecode) > 0:
            e = str(bytecode[0][0])
            t = [str(item) for item in bytecode[0][1:]]
            bytecode.pop(0)
            Decompiler.vg += 1
            
            if e in Decompiler.mapping:
                Decompiler.handle_operation(Decompiler.mapping[e], t)
            else:
                Decompiler.decompiled += f"// UNKNOWN_OPCODE {e} -> {', '.join(t)};\n"
            
            if Decompiler.mapping.get(e) == "CALL" and not Decompiler.found:
                for entry in Decompiler.potential:
                    if len(t) > 3 and entry["var"] == t[3]:
                        key_str = str(entry["key"]).replace(".", "_")
                        regex = rf"var var_{key_str} = (.*);"
                        match = re.search(regex, Decompiler.decompiled)
                        if match:
                            Decompiler.xorkey2 = match.group(1).replace(";", "")
                        Decompiler.found = True
                        break

        if Decompiler.round1 == 0:
            Decompiler.round1 += 1
            Decompiler.decompile_2()

    @staticmethod
    def decompile_2():
        matches = [m.group(2) for m in re.finditer(r"var\s+\w+\s*=\s*(['\"`])([\s\S]*?)\1", Decompiler.decompiled)]
        bytecode = max(matches, key=len, default="")
        if bytecode:
            decoded = json.loads(Decompiler.xS(base64.b64decode(bytecode).decode(), str(Decompiler.xorkey)))
            Decompiler.decompile(decoded)
        
        if Decompiler.round1 == 1:
            Decompiler.round1 += 1
            Decompiler.decompile_3()

    @staticmethod
    def decompile_3():
        matches = [m.group(2) for m in re.finditer(r"var\s+\w+\s*=\s*(['\"`])([\s\S]*?)\1", Decompiler.decompiled)]
        bytecode = next((s for s in matches if 60 <= len(s) <= 200), "")
        if bytecode:
            decoded = json.loads(Decompiler.xS(base64.b64decode(bytecode).decode(), str(Decompiler.xorkey)))
            Decompiler.decompile(decoded)
        Decompiler.remove_unused_variables()

    @staticmethod
    def decompile_vm(turnstile, token):
        Decompiler.start()
        Decompiler.decompiled = (
            "const { JSDOM } = require(\"jsdom\");\n"
            "const dom = new JSDOM(\"<!DOCTYPE html><p>Hello world</p>\", { url: \"https://chatgpt.com/\" });\n"
            "const window = dom.window;\n"
            "var mem = {};\n"
        )
        Decompiler.decompile(json.loads(Decompiler.xS(base64.b64decode(turnstile).decode(), str(token))))
        return Decompiler.decompiled
    
# NOTE dont mind this please i converted my JS decompiler to py using AI dont judge this please
</file>

<file path="wrapper/reverse/parse.py">
import esprima
import re

class Parser:

    @staticmethod
    def find_var_definition(var_name, start_line, code):
        code_lines = code.splitlines()

        relevant_code = '\n'.join(code_lines[:start_line - 1])

        sub_ast = esprima.parseScript(relevant_code, {'loc': True, 'range': True, 'tolerant': True})

        var_defs = {}

        def collect_var_defs(node, var_defs):
            if (node.type == 'VariableDeclarator' and 
                hasattr(node, 'id') and node.id and 
                hasattr(node, 'init') and node.init and 
                hasattr(node, 'loc') and node.loc):
                id_name = node.id.name if hasattr(node.id, 'name') else None
                if not id_name:
                    return
                abs_line = node.loc.start.line if hasattr(node.loc.start, 'line') else None
                if abs_line is None or abs_line >= start_line:
                    return
                if hasattr(node.init, 'range'):
                    value = relevant_code[node.init.range[0]:node.init.range[1]].strip()
                else:
                    value = str(node.init).strip() if node.init else ''
                if id_name not in var_defs:
                    var_defs[id_name] = []
                var_defs[id_name].append({'line': abs_line, 'value': value})

        def iterative_traverse(ast, visitor):
            if not ast:
                return
            stack = [ast]
            visited = set()
            max_stack_size = 10000

            while stack:
                if len(stack) > max_stack_size:
                    break
                node = stack.pop()
                node_id = id(node)
                if node_id in visited:
                    continue
                visited.add(node_id)

                visitor(node)
                for key in reversed(node.__dict__.keys()):
                    value = getattr(node, key, None)
                    if isinstance(value, list):
                        for item in reversed(value):
                            if isinstance(item, esprima.nodes.Node) and id(item) not in visited:
                                item._parent = node
                                stack.append(item)
                    elif isinstance(value, esprima.nodes.Node) and id(value) not in visited:
                        value._parent = node
                        stack.append(value)

        iterative_traverse(sub_ast, lambda n: collect_var_defs(n, var_defs))

        last_resolved = None
        def_line = None

        if var_name in var_defs:
            var_defs[var_name].sort(key=lambda x: x['line'], reverse=True)

            for defn in var_defs[var_name]:
                if 'btoa' not in defn['value'] and 'XOR_STR' not in defn['value'] and \
                'doubleXOR' not in defn['value'] and 'singlebtoa' not in defn['value']:
                    last_resolved = defn['value']
                    def_line = defn['line']
                    break

            if last_resolved:
                resolved_vars_cache = {}

                def resolve_var_recursive(expr, var_line):
                    try:
                        expr_ast = esprima.parseScript(expr, {'loc': True, 'range': True, 'tolerant': True})
                    except Exception:
                        return expr

                    vars_set = set()

                    def collect_identifiers(node):
                        if (hasattr(node, 'type') and node.type == 'Identifier' and 
                            hasattr(node, 'name')):
                            parent = getattr(node, '_parent', None)
                            if parent:
                                parent_type = parent.type if hasattr(parent, 'type') else None
                                if ((parent_type == 'MemberExpression' and 
                                    hasattr(parent, 'property') and parent.property == node and 
                                    not (hasattr(parent, 'computed') and parent.computed)) or
                                    (parent_type == 'ObjectProperty' and 
                                    hasattr(parent, 'key') and parent.key == node and 
                                    not (hasattr(parent, 'computed') and parent.computed)) or
                                    (parent_type == 'VariableDeclarator' and 
                                    hasattr(parent, 'id') and parent.id == node) or
                                    (parent_type == 'FunctionDeclaration' and 
                                    hasattr(parent, 'id') and parent.id == node) or
                                    (parent_type == 'FunctionExpression' and 
                                    hasattr(parent, 'id') and parent.id == node) or
                                    node.name == 'window'):
                                    return
                            vars_set.add(node.name)

                    def iterative_traverse_safe(ast, visitor):
                        if not ast:
                            return
                        stack = [ast]
                        visited = set()
                        while stack:
                            node = stack.pop()
                            node_id = id(node)
                            if node_id in visited:
                                continue
                            visited.add(node_id)
                            visitor(node)
                            for key in reversed(node.__dict__.keys()):
                                value = getattr(node, key, None)
                                if isinstance(value, list):
                                    for item in reversed(value):
                                        if isinstance(item, esprima.nodes.Node) and id(item) not in visited:
                                            item._parent = node
                                            stack.append(item)
                                elif isinstance(value, esprima.nodes.Node) and id(value) not in visited:
                                    value._parent = node
                                    stack.append(value)

                    iterative_traverse_safe(expr_ast, collect_identifiers)

                    if not vars_set:
                        return expr

                    for v in vars_set:
                        if v in resolved_vars_cache:
                            continue

                        def_value = v
                        if v in var_defs:
                            for defn in sorted(var_defs[v], key=lambda x: x['line'], reverse=True):
                                if defn['line'] < var_line and \
                                'btoa' not in defn['value'] and 'XOR_STR' not in defn['value'] and \
                                'doubleXOR' not in defn['value'] and 'singlebtoa' not in defn['value']:
                                    def_value = defn['value']
                                    break

                        resolved_vars_cache[v] = def_value
                        resolved_vars_cache[v] = resolve_var_recursive(def_value, var_line)

                    final_expr = expr
                    for k, v in resolved_vars_cache.items():
                        final_expr = re.sub(r'\b' + re.escape(k) + r'\b', str(v), final_expr)

                    return final_expr

                last_resolved = resolve_var_recursive(last_resolved, def_line)

                if last_resolved:
                    escaped_var_name = re.escape(var_name)

                    double_xor_pattern = re.compile(rf'XOR_STR\s*\(\s*{escaped_var_name}\s*,\s*{escaped_var_name}\s*\)')
                    xor_matches = double_xor_pattern.findall(code)

                    if xor_matches and len(xor_matches) >= 2:
                        last_resolved = f'doublexor({last_resolved})'
                    else:

                        usage_line_index = start_line - 1
                        search_start = max(0, usage_line_index - 10)
                        relevant_lines = '\n'.join(code_lines[search_start:usage_line_index + 1])

                        btoa_pattern = re.compile(rf'btoa\s*\(\s*""\s*\+\s*{escaped_var_name}\s*\)')
                        xor_var_pattern = re.compile(rf'XOR_STR\s*\(\s*{escaped_var_name}\s*,')

                        btoa_matches = btoa_pattern.findall(relevant_lines)
                        has_xor_var = bool(xor_var_pattern.search(relevant_lines))

                        if btoa_matches and len(btoa_matches) == 1 and not has_xor_var:
                            last_resolved = f'singlebtoa({last_resolved})'

        return last_resolved

    @staticmethod
    def parse_assigments(code):

            ast = esprima.parseScript(code, loc=True, jsx=True)

            stringify_calls = []

            def traverse_node(node):
                if isinstance(node, dict):
                    if node.get('type') == 'CallExpression':
                        callee = node.get('callee', {})
                        if (callee.get('type') == 'MemberExpression' and
                            callee.get('object', {}).get('name') == 'JSON' and
                            callee.get('property', {}).get('name') == 'stringify' and
                            node.get('arguments') and
                            node['arguments'][0]['type'] == 'Identifier'):
                            stringify_calls.append(node['arguments'][0]['name'])
                    for v in node.values():
                        traverse_node(v)
                elif isinstance(node, list):
                    for item in node:
                        traverse_node(item)

            traverse_node(ast.toDict())

            last_stringify_arg = stringify_calls[-1] if stringify_calls else None

            if not last_stringify_arg:
                return {}

            var_values = {}

            def traverse_vars(node):
                if isinstance(node, dict):
                    if node.get('type') == 'VariableDeclarator':
                        id_node = node.get('id', {})
                        init_node = node.get('init', {})
                        if (id_node.get('type') == 'Identifier' and
                            init_node and init_node.get('type') in ('Literal', 'NumericLiteral', 'StringLiteral')):
                            var_values[id_node['name']] = init_node.get('value')
                    for v in node.values():
                        traverse_vars(v)
                elif isinstance(node, list):
                    for item in node:
                        traverse_vars(item)

            traverse_vars(ast.toDict())

            assignments = {}

            def traverse_assignments(node):
                if isinstance(node, dict):
                    if node.get('type') == 'AssignmentExpression':
                        left = node.get('left', {})
                        right = node.get('right', {})
                        if (left.get('type') == 'MemberExpression' and
                            left.get('object', {}).get('type') == 'Identifier' and
                            left.get('object', {}).get('name') == last_stringify_arg and
                            left.get('property', {}).get('type') == 'Identifier' and
                            right.get('type') == 'Identifier' and
                            node.get('loc')):
                            key_var = left['property']['name']
                            value = right['name']
                            key = var_values.get(key_var, key_var)
                            resolved_value = Parser.find_var_definition(value, node['loc']['start']['line'], code) or value
                            assignments[key] = resolved_value
                    for v in node.values():
                        traverse_assignments(v)
                elif isinstance(node, list):
                    for item in node:
                        traverse_assignments(item)

            traverse_assignments(ast.toDict())

            return assignments
    @staticmethod
    def get_xor_key(js_code: str):
        
        parsed = esprima.parseScript(js_code, tolerant=True)
        
        last_xor_call = None
        second_arg_node = None

        for node in parsed.body:
            if node.type == 'VariableDeclaration':
                for decl in node.declarations:
                    if decl.init and decl.init.type == 'CallExpression':
                        call = decl.init
                        if call.callee.type == 'Identifier' and call.callee.name == 'XOR_STR':
                            last_xor_call = call
                            second_arg_node = call.arguments[1]

        if not last_xor_call:
            return None

        if second_arg_node.type == 'Identifier':
            var_name = second_arg_node.name
        elif second_arg_node.type == 'Literal':
            return second_arg_node.value
        else:
            return None

        def find_value(nodes, name):
            for node in nodes:
                if node.type == 'VariableDeclaration':
                    for decl in node.declarations:
                        if decl.id.name == name and decl.init.type == 'Literal':
                            return decl.init.value
                elif node.type == 'ExpressionStatement' and node.expression.type == 'AssignmentExpression':
                    expr = node.expression
                    if expr.left.type == 'Identifier' and expr.left.name == name and expr.right.type == 'Literal':
                        return expr.right.value
            return None

        return find_value(parsed.body, var_name)
    
    @staticmethod
    def parse_keys(decompiled_code: str) -> tuple[str, dict]:
        
        assignments: dict = Parser.parse_assigments(decompiled_code)
        xor_key: str = Parser.get_xor_key(decompiled_code)
        
        parsed_keys: dict = {}
        randomindex = 1
        for key, value in assignments.items():
            key = str(key)
            if value.startswith("Array") and "location" not in value:
                numbers = value.split(') : ')[1].split(" + ")
                num1 = float(numbers[0])
                num2 = float(numbers[1])
                parsed_keys[key] = str(float(num1 + num2))
            elif "location" in value:
                parsed_keys[key] = "location"
            elif "cfIpLongitude" in value:
                parsed_keys[key] = "ipinfo"
            elif "maxTouchPoints" in value:
                parsed_keys[key] = "vendor"
            elif "history" in value:
                parsed_keys[key] = "history"
            elif 'window["Object"]["keys"]' in value:
                parsed_keys[key] = "localstorage"
            elif 'createElement' in value:
                parsed_keys[key] = "element"
            elif value.isdigit():
                parsed_keys[key] = value
            elif "random" in value:
                parsed_keys[key] = "random_" + str(randomindex)
                randomindex += 1
            elif "doublexor" in value:
                parsed_keys[key] = value
            elif "singlebtoa" in value:
                parsed_keys[key] = value

        return xor_key, parsed_keys
    
# NOTE dont mind this please i converted my babel parser from JS to py using AI i know its shit but i didnt wanna exec js code or smt
</file>

<file path="wrapper/reverse/vm.py">
from json        import dumps, loads
from base64      import b64decode, b64encode
from random      import randint, random
from .decompiler import Decompiler
from .parse      import Parser


class VM:
    
    html_object: str = dumps({"x":0,"y":1219,"width":37.8125,"height":30,"top":1219,"right":37.8125,"bottom":1249,"left":0}, separators=(',', ':'))

    @staticmethod
    def xor(e, t):
        t = str(t)
        e = str(e)
        n = ""

        for r in range(len(e)):
            n += chr(ord(e[r]) ^ ord(t[r % len(t)]))

        return n
    
    @staticmethod
    def get_turnstile(turnstile: str, token: str, ip_info: str) -> str:
        
        decompiled: str = Decompiler.decompile_vm(turnstile, token)
        
        xor_key, keys = Parser.parse_keys(decompiled)

        payload: dict = {}
    
        for key, value in keys.items():
            try:
                value = float(value)
            except:
                ...
            
            if isinstance(value, float):
                payload[key] = b64encode(VM.xor(str(value), xor_key).encode("utf-8")).decode("utf-8")
                
            elif "singlebtoa" in value:
                payload[key] = b64encode(value.split("singlebtoa(")[1].split(")")[0].encode("utf-8")).decode("utf-8")
            
            elif "doublexor" in value:
                number: str = value.split("doublexor(")[1].split(")")[0]
                value_1: str = b64encode(VM.xor(number, number).encode("utf-8")).decode("utf-8")
                value_2: str = b64encode(VM.xor(value_1, value_1).encode("utf-8")).decode("utf-8")
                payload[key] = b64encode(value_2.encode("utf-8")).decode("utf-8")
            
            elif "ipinfo" in value:
                payload[key] = b64encode(VM.xor(ip_info, xor_key).encode("utf-8")).decode("utf-8")
            
            elif "element" in value:
                payload[key] = b64encode(VM.xor(VM.html_object, xor_key).encode()).decode()

            elif "location" in value:
                location: str = 'https://chatgpt.com/'
                payload[key] = b64encode(VM.xor(location, xor_key).encode("utf-8")).decode("utf-8")
            
            elif "random_1" in value:
                random_value: float = random()
                payload[key] = b64encode(VM.xor(str(random_value), str(random_value)).encode("utf-8")).decode("utf-8")
            
            elif "random_2" in value:
                payload[key] = random()
                
            elif "vendor" in value:
                vendor_info: str = '["Google Inc.","Win32",8,0]'
                payload[key] = b64encode(VM.xor(vendor_info, xor_key).encode("utf-8")).decode("utf-8")
            
            elif "localstorage" in value:
                payload[key] = b64encode(VM.xor('oai/apps/hasDismissedTeamsNoAuthUpsell,oai/apps/lastSeenNoAuthTrialsBannerAt,oai-did,oai/apps/noAuthGoUpsellModalDismissed,oai/apps/hasDismissedBusinessFreeTrialUpsellModal,oai/apps/capExpiresAt,statsig.session_id.1792610830,oai/apps/hasSeenNoAuthImagegenNux,oai/apps/lastPageLoadDate,client-correlated-secret,statsig.stable_id.1792610830,oai/apps/debugSettings,oai/apps/hasDismissedPlusFreeTrialUpsellModal,oai/apps/tatertotInContextUpsellBannerV2,search.attributions-settings', xor_key).encode("utf-8")).decode("utf-8")
            
            elif "history" in value:
                payload[key] = b64encode(VM.xor(str(randint(1, 5)), xor_key).encode()).decode()
            
            else:
                print(f"UNKNOWN ITEM WTFFFF {key},{value}")
        
        turnstile_token: str = b64encode(VM.xor(dumps(payload, separators=(',', ':')), xor_key).encode("utf-8")).decode("utf-8")
        
        return turnstile_token
</file>

<file path="wrapper/runtime.py">
from typing    import Callable, Any, Optional, Type
from .logger   import Log
from functools import wraps


class Run:
    """
    Class to handle runtime
    """
    
    @staticmethod
    def Error(func: Callable[..., Any]) -> Callable[..., Any]:
        """
        Error function to catch errors
        
        @param func: The function to wrap.
        @return:     Custom error message
        """
        @wraps(func)
        def wrapper(*args: Any, **kwargs: Any) -> Any:
            try:
                return func(*args, **kwargs)
            except Exception as e:
                Run.handle_error(e)
                return None 
        return wrapper

    @staticmethod
    def handle_error(exception: Exception) -> Optional[None]:
        """
        Handling an error
        
        @param exception: Exception that occured
        """
        Log.Error(f"Error occurred: {exception}")
        exit()
        
class Utils:
    
    @staticmethod
    def between(
        main_text: Optional[str],
        value_1: Optional[str],
        value_2: Optional[str],
        ) -> Type[str]:
        return main_text.split(value_1)[1].split(value_2)[0]
</file>

</files>
