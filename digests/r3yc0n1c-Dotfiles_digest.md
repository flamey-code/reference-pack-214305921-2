This file is a merged representation of the entire codebase, combined into a single document by Repomix.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)

# Directory Structure
```
.tmux.conf
.vimrc
Antigravity/install-open-with-antigravity.reg
Antigravity/uninstall-open-with-antigravity.reg
chatgpt/delete_mem.py
Claude/README.md
cleaner.sh
ginit.vim
init.vim
Microsoft.PowerShell_profile.ps1
README.md
Sublime-Text-3/add_date.py
Sublime-Text-3/newPython3.sublime-build
Sublime-Text-3/python_about.sublime-snippet
Sublime-Text-3/python_shebang.sublime-snippet
Sublime-Text-3/PythonREPL.sublime-build
Sublime-Text-3/README.md
Sublime-Text-3/terminus_python.sublime-build
Sublime-Text-3/terminus_repl.py
Windows_Terminal/Microsoft.PowerShell_profile.ps1
Windows_Terminal/README.md
Windows_Terminal/settings.json
```

# Files

## File: .tmux.conf
````ini
set -g activity-action other
set -g assume-paste-time 1
set -g base-index 0
set -g bell-action any
#set -g default-command 
set -g default-shell /usr/bin/zsh
set -g default-size 80x24
set -g destroy-unattached off
set -g detach-on-destroy on
set -g display-panes-active-colour red
set -g display-panes-colour blue
set -g display-panes-time 1000
set -g display-time 750
set -g history-limit 2000
set -g key-table root
set -g lock-after-time 0
set -g lock-command "lock -np"
set -g message-command-style fg=yellow,bg=black
set -g message-style fg=black,bg=yellow
set -g mouse on
set -g prefix C-a
unbind-key C-b
bind-key C-a send-prefix
#set -g prefix2 Invalid#1fff00000000
set -g renumber-windows off
set -g repeat-time 500
set -g set-titles off
set -g set-titles-string "#S:#I:#W - \"#T\" #{session_alerts}"
set -g silence-action other
set -g status on
set -g status-bg green
set -g status-fg black
set -g status-format[0] "#[align=left range=left #{status-left-style}]#[push-default]#{T;=/#{status-left-length}:status-left}#[pop-default]#[norange default]#[list=on align=#{status-justify}]#[list=left-marker]<#[list=right-marker]>#[list=on]#{W:#[range=window|#{window_index} #{window-status-style}#{?#{&&:#{window_last_flag},#{!=:#{window-status-last-style},default}}, #{window-status-last-style},}#{?#{&&:#{window_bell_flag},#{!=:#{window-status-bell-style},default}}, #{window-status-bell-style},#{?#{&&:#{||:#{window_activity_flag},#{window_silence_flag}},#{!=:#{window-status-activity-style},default}}, #{window-status-activity-style},}}]#[push-default]#{T:window-status-format}#[pop-default]#[norange default]#{?window_end_flag,,#{window-status-separator}},#[range=window|#{window_index} list=focus #{?#{!=:#{window-status-current-style},default},#{window-status-current-style},#{window-status-style}}#{?#{&&:#{window_last_flag},#{!=:#{window-status-last-style},default}}, #{window-status-last-style},}#{?#{&&:#{window_bell_flag},#{!=:#{window-status-bell-style},default}}, #{window-status-bell-style},#{?#{&&:#{||:#{window_activity_flag},#{window_silence_flag}},#{!=:#{window-status-activity-style},default}}, #{window-status-activity-style},}}]#[push-default]#{T:window-status-current-format}#[pop-default]#[norange list=on default]#{?window_end_flag,,#{window-status-separator}}}#[nolist align=right range=right #{status-right-style}]#[push-default]#{T;=/#{status-right-length}:status-right}#[pop-default]#[norange default]"
set -g status-format[1] "#[align=centre]#{P:#{?pane_active,#[reverse],}#{pane_index}[#{pane_width}x#{pane_height}]#[default] }"
set -g status-interval 15
set -g status-justify left
set -g status-keys emacs
set -g status-left "[#S] "
set -g status-left-length 10
set -g status-left-style default
set -g status-position bottom
set -g status-right "#{?window_bigger,[#{window_offset_x}#,#{window_offset_y}] ,}\"#{=21:pane_title}\" %H:%M %d-%b-%y"
set -g status-right-length 40
set -g status-right-style default
set -g status-style fg=black,bg=green
set -g update-environment[0] DISPLAY
set -g update-environment[1] KRB5CCNAME
set -g update-environment[2] SSH_ASKPASS
set -g update-environment[3] SSH_AUTH_SOCK
set -g update-environment[4] SSH_AGENT_PID
set -g update-environment[5] SSH_CONNECTION
set -g update-environment[6] WINDOWID
set -g update-environment[7] XAUTHORITY
set -g visual-activity off
set -g visual-bell off
set -g visual-silence off
set -g word-separators " "
````

## File: .vimrc
````
syntax on
set mouse=a

set number " nu
set tabstop=4 " ts
set shiftwidth=4 " sw
set softtabstop=4 " sts
set autoindent " ai
set smartindent " si
set wrap
set hlsearch
set ignorecase
set smartcase
set smarttab
set magic
set showmatch
set cursorline

filetype indent on


" run code
augroup compileandrun
    autocmd!
    autocmd filetype python nnoremap <F5> :w <bar> :!py % <cr>
    autocmd filetype cpp nnoremap <F5> :w <bar> !g++ % -o %< && ./%< <cr>
augroup END
````

## File: Antigravity/install-open-with-antigravity.reg
````
Windows Registry Editor Version 5.00

[HKEY_CURRENT_USER\Software\Classes\*\shell\Antigravity]
@="Open with Antigravity"
"Icon"="C:\\Users\\rey\\AppData\\Local\\Programs\\Antigravity\\Antigravity.exe"

[HKEY_CURRENT_USER\Software\Classes\*\shell\Antigravity\command]
@="\"C:\\Users\\rey\\AppData\\Local\\Programs\\Antigravity\\Antigravity.exe\" \"%1\""

[HKEY_CURRENT_USER\Software\Classes\Directory\shell\Antigravity]
@="Open with Antigravity"
"Icon"="C:\\Users\\rey\\AppData\\Local\\Programs\\Antigravity\\Antigravity.exe"

[HKEY_CURRENT_USER\Software\Classes\Directory\shell\Antigravity\command]
@="\"C:\\Users\\rey\\AppData\\Local\\Programs\\Antigravity\\Antigravity.exe\" \"%V\""

[HKEY_CURRENT_USER\Software\Classes\Directory\Background\shell\Antigravity]
@="Open with Antigravity"
"Icon"="C:\\Users\\rey\\AppData\\Local\\Programs\\Antigravity\\Antigravity.exe"

[HKEY_CURRENT_USER\Software\Classes\Directory\Background\shell\Antigravity\command]
@="\"C:\\Users\\rey\\AppData\\Local\\Programs\\Antigravity\\Antigravity.exe\" \"%V\""
````

## File: Antigravity/uninstall-open-with-antigravity.reg
````
Windows Registry Editor Version 5.00

[-HKEY_CURRENT_USER\Software\Classes\*\shell\Antigravity]

[-HKEY_CURRENT_USER\Software\Classes\Directory\shell\Antigravity]

[-HKEY_CURRENT_USER\Software\Classes\Directory\Background\shell\Antigravity]
````

## File: chatgpt/delete_mem.py
````python
'''
Author: r3yc0n1c
Date: 05-05-2026
Context: Clear old ChatGPT Memory to free model context
'''

import requests
from datetime import datetime, timedelta, date

BASE_URL = "https://chatgpt.com/backend-api/memories"  # replace with real API
TOKEN = "<YOUR-BEARER-TOKEN>"

HEADERS = {
    "Authorization": f"Bearer {TOKEN}",
    "Content-Type": "application/json"
}

def fetch_memories():
    URL = f"{BASE_URL}?exclusive_to_gizmo=false&include_memory_entries=true"
    resp = requests.get(URL, headers=HEADERS)
    resp.raise_for_status()
    return resp.json()

def delete_memory(memory_id):
    url = f"{BASE_URL}/{memory_id}"
    resp = requests.delete(url, headers=HEADERS)
    if resp.status_code not in (200, 204):
        print(f"Failed to delete {memory_id}: {resp.status_code}")

def is_qualified_for_delete(updated_at, days_threshold=30):
    updated = datetime.fromisoformat(updated_at).date()
    cutoff = date.today() - timedelta(days=days_threshold)

    print(f"Memory updated at {updated}, cutoff is {cutoff}")
    return updated < cutoff

def main():
    data = fetch_memories()

    # adjust depending on API structure
    memories = data.get("memories", [])
    # test
    # print(memories)
    # memories = [{'id': '102b7e6c-bc80-41a0-b929-5a139d897cd1', 'content': 'fish cat top of the house.', 'updated_at': '2024-09-01', 'gizmo_id': None, 'status': 'warm', 'conversation_id': None, 'created_timestamp': None, 'last_updated': None, 'labels': None}]

    for mem in memories:
        updated_at = mem.get("updated_at")
        memory_id = mem.get("id")

        if not updated_at or not memory_id:
            continue

        if is_qualified_for_delete(updated_at, days_threshold=365):
            print(f"Deleting {memory_id} (updated {updated_at})")
            delete_memory(memory_id)

if __name__ == "__main__":
    main()
````

## File: Claude/README.md
````markdown
## Setup Claude Code (Free with Gemini models)

### 1. Install Claude Code and verrify
```powershell 
npm install -g @anthropic-ai/claude-code

# verify
claude --version
```

### 2. Get Gemini API KEY

- Go to https://aistudio.google.com/ and click on "Get API Key" in the bottom left corner
- Save it somewhere

### 3. Cook the Proxy

- Install LiteLLM
```powershell 
pip install 'litellm[proxy]'
```

- Create `litellm_config.yaml` in your home folder
```powershell
touch C:\Users\<YOUR-USERNAME>\litellm_config.yaml
```
- Put these configs and save the file
```
model_list:
  - model_name: gemini-2.5-flash
    litellm_params:
      model: gemini/gemini-2.5-flash
      api_key: YOUR-GEMINI-KEY

litellm_settings:
  drop_params: true
```

- Start the proxy
```powershell
cd ~
litellm --config litellm_config.yaml --port 4000
```

- Claude Code flow `Claude Code -> Anthropic format -> LiteLLM -> OpenAI format -> Gemini`


### 4. Setup ENV variables

Windows
```powershell
setx ANTHROPIC_API_KEY "AIza-YOUR-GEMINI-KEY-HERE"
setx ANTHROPIC_AUTH_TOKEN "AIza-YOUR-GEMINI-KEY-HERE"
setx ANTHROPIC_BASE_URL "http://localhost:4000"
setx ANTHROPIC_MODEL "gemini-2.5-flash"
```

Linux/MacOS
```sh
ANTHROPIC_API_KEY=YOUR-GEMINI-KEY-HERE
ANTHROPIC_AUTH_TOKEN=YOUR-GEMINI-KEY-HERE
ANTHROPIC_BASE_URL=http://localhost:4000
ANTHROPIC_MODEL=gemini-2.5-flash
```

### 5. Refresh `env` and start Claude Code
```powershell
cd your_project
claude .

# When it asks for "Do you want to use this API key?"
# Select "No" and enter
```
````

## File: cleaner.sh
````bash
#!/bin/bash

find . -type f -executable -not -path '*/.git/*' -not -name "*.sh" -delete
````

## File: ginit.vim
````vim
" save plugins in this directory
call plug#begin('~/AppData/Local/nvim/plugged')

" below are some vim plugins i like to use...

    Plug 'junegunn/fzf', { 'do': { -> fzf#install() } }
    Plug 'junegunn/fzf.vim'
    Plug 'sheerun/vim-polyglot'
    Plug 'Raimondi/delimitMate'
    Plug 'vim-airline/vim-airline'
    Plug 'vim-syntastic/syntastic'
    Plug 'vim-airline/vim-airline-themes'
    Plug 'rafi/awesome-vim-colorschemes'
    Plug 'preservim/nerdcommenter'
    Plug 'ervandew/supertab'
    Plug 'preservim/nerdtree' " Tree explorer
    Plug 'msanders/snipmate.vim' " Quick snippets
    Plug 'neoclide/coc.nvim', {'branch': 'release'} " Auto Completion    

" coc config
"
" :CocInstall coc-clangd coc-python coc-html coc-css coc-json coc-tsserver
" :CocInstall coc-snippets coc-emmet coc-highlight coc-prettier coc-pairs
" :CocInstall coc-spell-checker coc-eslint

call plug#end()

" ----- GUI Customization START -----

" Enable Mouse
set mouse=a

" Set Editor Font
if exists(':GuiFont')
    " Use GuiFont! to ignore font errors
    GuiFont! Fira Code:h12
endif

" Disable GUI Tabline
if exists(':GuiTabline')
    GuiTabline 0
endif

" Disable GUI Popupmenu
if exists(':GuiPopupmenu')
    GuiPopupmenu 0
endif

" Enable GUI ScrollBar
if exists(':GuiScrollBar')
    GuiScrollBar 1
endif

" Right Click Context Menu (Copy-Cut-Paste)
nnoremap <silent><RightMouse> :call GuiShowContextMenu()<CR>
inoremap <silent><RightMouse> <Esc>:call GuiShowContextMenu()<CR>
xnoremap <silent><RightMouse> :call GuiShowContextMenu()<CR>gv
snoremap <silent><RightMouse> <C-G>:call GuiShowContextMenu()<CR>gv

" ----- GUI Customization END -----

filetype plugin on

let mapleader = "-"
let maplocalleader = "\\"

"split navigations
nnoremap <C-J> <C-W><C-J>
nnoremap <C-K> <C-W><C-K>
nnoremap <C-L> <C-W><C-L>
nnoremap <C-H> <C-W><C-H>

set splitbelow
set splitright

" Nerdtree settings
nnoremap <leader>n :NERDTreeFocus<CR>
nnoremap <C-n> :NERDTree<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
nnoremap <C-f> :NERDTreeFind<CR>
" let g:NERDTreeWinPos = "right"

" Enable folding
set foldmethod=indent
set foldlevel=99
"Enable folding with the spacebar
nnoremap <space> za

" open files with ctrl-p
nnoremap <c-p> :Files<cr>

au BufNewFile,BufRead *.py,*.java,*.cpp,*.c,*.cs,*.rkt,*.h,*.html
    \ set tabstop=4 |
    \ set softtabstop=4 |
    \ set shiftwidth=4 |
    \ set textwidth=120 |
    \ set expandtab |
    \ set autoindent |
    \ set fileformat=unix |

set encoding=UTF-8

syntax on

" air-line
let g:airline_powerline_fonts = 1
let g:airline_theme = 'luna'
let g:airline#extensions#tabline#enabled = 1

if !exists('g:airline_symbols')
    let g:airline_symbols = {}
endif

" unicode symbols
let g:airline_left_sep = '»'
let g:airline_left_sep = '▶'
let g:airline_right_sep = '«'
let g:airline_right_sep = '◀'
let g:airline_symbols.linenr = '␊'
let g:airline_symbols.linenr = '␤'
let g:airline_symbols.linenr = '¶'
let g:airline_symbols.branch = '⎇'
let g:airline_symbols.paste = 'ρ'
let g:airline_symbols.paste = 'Þ'
let g:airline_symbols.paste = '∥'
let g:airline_symbols.whitespace = 'Ξ'

" airline symbols
let g:airline_left_sep = ''
let g:airline_left_alt_sep = ''
let g:airline_right_sep = ''
let g:airline_right_alt_sep = ''
let g:airline_symbols.branch = ''
let g:airline_symbols.readonly = ''
let g:airline_symbols.linenr = ''

highlight Comment cterm=italic gui=italic

set laststatus=2
" set showtabline=2

" true colours
set background=dark
set t_Co=256

if (has("nvim"))
  let $NVIM_TUI_ENABLE_TRUE_COLOR=1
endif

if (has("termguicolors"))
  set termguicolors
endif

colorscheme jellybeans

set nu rnu " relative line numbering
set clipboard=unnamed " public copy/paste register
set ruler
set showcmd
set noswapfile " doesn't create swap files
set noshowmode
set shortmess+=c
set omnifunc=syntaxcomplete#Complete

set backspace=indent,eol,start " let backspace delete over lines
set autoindent " enable auto indentation of lines
set smartindent " allow vim to best-effort guess the indentation
set pastetoggle=<F2> " enable paste mode

set wildmenu "graphical auto complete menu
set lazyredraw "redraws the screne when it needs to
set showmatch "highlights matching brackets
set incsearch "search as characters are entered
set hlsearch "highlights matching searches

"clears highlights
nnoremap // :noh<return>
" moves current line down or up
nnoremap <leader>- ddp
nnoremap <leader>_ ddkP
" open vimrc in vertical split
nnoremap <leader>ev :vsplit $MYVIMRC<cr>
" update changes into current buffer
nnoremap <leader>sv :source $MYVIMRC<cr>
" enable or disable line wrapping in current buffer
nnoremap <buffer> <localleader>w :set wrap!<cr>

" c++11 support in syntastic
let g:syntastic_cpp_compiler = 'clang++'
let g:syntastic_cpp_compiler_options = ' -std=c++11'

" run code
augroup compileandrun
    autocmd!
    autocmd filetype python nnoremap <f5> :w <bar> :!py % <cr>
"    autocmd filetype cpp nnoremap <f5> :w <bar> !g++ -std=c++11 % <cr> :vnew <bar> :te "a.exe" <cr><cr>
"    autocmd filetype cpp nnoremap <f6> :vnew <bar> :te "a.exe" <cr>
    autocmd filetype cpp nnoremap <f5> :w <bar> !g++ % -o %< && %< <cr>
    autocmd filetype c nnoremap <f5> :w <bar> !gcc % -o %< && %< <cr>
    autocmd filetype java nnoremap <f5> :w <bar> !javac % && java %:r <cr>
augroup END
````

## File: init.vim
````vim
" save plugins in this directory
call plug#begin('~/AppData/Local/nvim/plugged')

" list of vim plugins...

    Plug 'junegunn/fzf', { 'do': { -> fzf#install() } }
    Plug 'junegunn/fzf.vim'
    Plug 'sheerun/vim-polyglot'
    Plug 'Raimondi/delimitMate'
    Plug 'vim-airline/vim-airline'
    Plug 'vim-syntastic/syntastic'
    Plug 'vim-airline/vim-airline-themes'
    Plug 'rafi/awesome-vim-colorschemes'
    Plug 'preservim/nerdcommenter'
    Plug 'ervandew/supertab'
    Plug 'preservim/nerdtree' " Tree explorer
    Plug 'msanders/snipmate.vim' " Quick snippets
    Plug 'neoclide/coc.nvim', {'branch': 'release'} " Auto Completion    

" coc config
"
" :CocInstall coc-clangd coc-python coc-html coc-css coc-json coc-tsserver
" :CocInstall coc-snippets coc-emmet coc-highlight coc-prettier coc-pairs
" :CocInstall coc-spell-checker coc-eslint

call plug#end()

" Enable Mouse
set mouse=a

filetype plugin on

"let mapleader = "-"
"let maplocalleader = "\\"

"split navigations
nnoremap <C-J> <C-W><C-J>
nnoremap <C-K> <C-W><C-K>
nnoremap <C-L> <C-W><C-L>
nnoremap <C-H> <C-W><C-H>

set splitbelow
set splitright

" Nerdtree settings
nnoremap <leader>n :NERDTreeFocus<CR>
nnoremap <C-n> :NERDTree<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
" nnoremap <C-f> :NERDTreeFind<CR>
" let g:NERDTreeWinPos = "right"

" Enable folding
set foldmethod=indent
set foldlevel=99
"Enable folding with the spacebar
nnoremap <space> za

" open files with ctrl-p
nnoremap <c-p> :Files<cr>

au BufNewFile,BufRead *.py,*.java,*.cpp,*.c,*.cs,*.rkt,*.h,*.html
    \ set tabstop=4 |
    \ set softtabstop=4 |
    \ set shiftwidth=4 |
    \ set textwidth=120 |
    \ set expandtab |
    \ set autoindent |
    \ set fileformat=unix |

set encoding=utf-8

syntax on

" air-line
let g:airline_powerline_fonts = 1
let g:airline_theme = 'jellybeans'
let g:airline#extensions#tabline#enabled = 1

if !exists('g:airline_symbols')
    let g:airline_symbols = {}
endif

" unicode symbols
let g:airline_left_sep = '»'
let g:airline_left_sep = '▶'
let g:airline_right_sep = '«'
let g:airline_right_sep = '◀'
let g:airline_symbols.linenr = '␊'
let g:airline_symbols.linenr = '␤'
let g:airline_symbols.linenr = '¶'
let g:airline_symbols.branch = '⎇'
let g:airline_symbols.paste = 'ρ'
let g:airline_symbols.paste = 'Þ'
let g:airline_symbols.paste = '∥'
let g:airline_symbols.whitespace = 'Ξ'

" airline symbols
let g:airline_left_sep = ''
let g:airline_left_alt_sep = ''
let g:airline_right_sep = ''
let g:airline_right_alt_sep = ''
let g:airline_symbols.branch = ''
let g:airline_symbols.readonly = ''
let g:airline_symbols.linenr = ''

highlight Comment cterm=italic gui=italic

set laststatus=2
" set showtabline=2

" true colours
set background=dark
set t_Co=256

if (has("nvim"))
  let $NVIM_TUI_ENABLE_TRUE_COLOR=1
endif

if (has("termguicolors"))
  set termguicolors
endif

colorscheme jellybeans

set nu rnu " relative line numbering
set clipboard=unnamed " public copy/paste register
set ruler
set showcmd
set noswapfile " doesn't create swap files
set noshowmode
set shortmess+=c
set omnifunc=syntaxcomplete#Complete

set backspace=indent,eol,start " let backspace delete over lines
set autoindent " enable auto indentation of lines
set smartindent " allow vim to best-effort guess the indentation
set pastetoggle=<F2> " enable paste mode

set wildmenu "graphical auto complete menu
set lazyredraw "redraws the screne when it needs to
set showmatch "highlights matching brackets
set incsearch "search as characters are entered
set hlsearch "highlights matching searches

"clears highlights
nnoremap // :noh<return>
" moves current line down or up
nnoremap <leader>- ddp
nnoremap <leader>_ ddkP
" open vimrc in vertical split
nnoremap <leader>ev :vsplit $MYVIMRC<cr>
" update changes into current buffer
nnoremap <leader>sv :source $MYVIMRC<cr>
" enable or disable line wrapping in current buffer
nnoremap <buffer> <localleader>w :set wrap!<cr>

" c++11 support in syntastic
let g:syntastic_cpp_compiler = 'clang++'
let g:syntastic_cpp_compiler_options = ' -std=c++11'

" run code
augroup compileandrun
    autocmd!
    autocmd filetype python nnoremap <F5> :w <bar> :!py % <cr>
"    autocmd filetype cpp nnoremap <f5> :w <bar> !g++ -std=c++11 % <cr> :vnew <bar> :te "a.exe" <cr><cr>
"    autocmd filetype cpp nnoremap <F6> :vnew <bar> :te "a.exe" <cr>
    autocmd filetype cpp nnoremap <F5> :w <bar> !g++ % -o %< && %< <cr>
    autocmd filetype c nnoremap <F6> :w <bar> !gcc % <cr>
    autocmd filetype c nnoremap <F5> :w <bar> !gcc % <cr> :vnew <bar> :te "a.exe" <cr><cr>
    autocmd filetype c nnoremap <F7> :vnew <bar> :te "a.exe" <cr>
    "autocmd filetype c nnoremap <F5> :w <bar> !gcc % -o %< && %< <cr>
    autocmd filetype java nnoremap <F5> :w <bar> !javac % && java %:r <cr>
augroup END
````

## File: Microsoft.PowerShell_profile.ps1
````powershell
# Custom PS Profile inspired from theme - robbyrussell
# Author: r3yc0n1c

# clear startup text
clear

function Write-BranchName () {
    try {
        $branch = git rev-parse --abbrev-ref HEAD

        if ($branch -eq "HEAD") {
            # we're probably in detached HEAD state, so print the SHA
            $branch = git rev-parse --short HEAD
            Write-Host " git:($branch)" -ForegroundColor "red"
        }
        else {
            # we're on an actual branch, so print it
            Write-Host " git:($branch)" -ForegroundColor "blue"
        }
    } catch {
        # we'll end up here if we're in a newly initiated git repo
        Write-Host " git:(no branches yet)" -ForegroundColor "yellow"
    }
}

function prompt {
    $base = "PS "    
    # $path = "($(Get-Location)) "
    $path = "($(((Get-Location) | Get-Item).Name))"
    # $time = "($(Get-Date))"
    # $dt = "($(Get-Date -Format "dddd MM/dd/yyyy HH:mm:ss"))"
    # $userPrompt = "$("(>^-^)>" * ($nestedPromptLevel + 1)) "
    $userPrompt = "(>"

    Write-Host $base -NoNewline -ForegroundColor "red"

    if (Test-Path .git) {
        Write-Host $path -NoNewline -ForegroundColor "green"
        # Write-Host $dt -ForegroundColor "yellow"
        Write-BranchName
    }
    else {
        # we're not in a repo so don't bother displaying branch name/sha
        Write-Host $path -ForegroundColor "green"
        # Write-Host $dt -ForegroundColor "yellow"
    }

    Write-Host $userPrompt -NoNewline -ForegroundColor "red"
    return " "
}

# MY ALIASES
function launchIPython { python -m IPython }
# function execTime { (Measure-Command { echo $args | Out-Default }).ToString() }
# function execTime { Measure-Command {cmd /c $args | Out-Host} }
function launchBash { & 'C:\Program Files\Git\bin\sh.exe' --login }

Set-Alias ipy launchIPython
Set-Alias sh launchBash
# Set-Alias time execTime
````

## File: README.md
````markdown
# Config Files
This repository contains all the special Configuration Files created, tested and implemented by me on my local machine(s).

## :warning: Caution!
These files may or may not work for you! Edit them before using.
````

## File: Sublime-Text-3/add_date.py
````python
import datetime, getpass
import sublime, sublime_plugin
class AddDateCommand(sublime_plugin.TextCommand):
    def run(self, edit):
        self.view.run_command("insert_snippet", { "contents": "%s" %  datetime.date.today().strftime("%d %B %Y (%A)") } )

class AddTimeCommand(sublime_plugin.TextCommand):
    def run(self, edit):
        self.view.run_command("insert_snippet", { "contents": "%s" %  datetime.datetime.now().strftime("%H:%M") } )
````

## File: Sublime-Text-3/newPython3.sublime-build
````
{
 "cmd": ["/usr/bin/python3", "-u", "$file"],
 "file_regex": "^[ ]File \"(...?)\", line ([0-9]*)",
 "selector": "source.python"
}
````

## File: Sublime-Text-3/python_about.sublime-snippet
````
<snippet>
	<content><![CDATA[
"""
Author:       YOU
Date:         ${1:edit_me}
Description:  ${2:edit_me}
"""
]]></content>
	<!-- Optional: Set a tabTrigger to define how to trigger the snippet -->
	<tabTrigger>about</tabTrigger>
	<!-- Optional: Set a scope to limit where the snippet will trigger -->
	<scope>source.python</scope>
	<description>about the script</description>
</snippet>
````

## File: Sublime-Text-3/python_shebang.sublime-snippet
````
<snippet>
	<content><![CDATA[
#!/usr/bin/env python3
]]></content>
	<!-- Optional: Set a tabTrigger to define how to trigger the snippet -->
	<tabTrigger>shebang</tabTrigger>
	<!-- Optional: Set a scope to limit where the snippet will trigger -->
	<scope>source.python</scope>
	<description>python3 shebang</description>
</snippet>
````

## File: Sublime-Text-3/PythonREPL.sublime-build
````
{
    "target": "terminus_open",
    "title": "Python Output",
    "tag": "python",
    "auto_close": false,

    // "shell_cmd": "python -u -i \"$file\"",
    "cmd": ["/usr/bin/python3", "-u", "$file"],
    "file_regex": "^[ ]*File \"(...*?)\", line ([0-9]*)",
    "selector": "source.python",

    "pre_window_hooks": [
    	["window_focus", {"store": true}],
    	["close_terminus_view_by_title", {"title": "Python Output"}],
    	["window_focus", {"store": false}]
    ],
    "post_window_hooks": [
        ["carry_file_to_pane", {"direction": "right"}],
        ["window_focus", {"store": false}]
    ],

    "env": {"PYTHONIOENCODING": "utf-8"}
}
````

## File: Sublime-Text-3/README.md
````markdown
## Config files for Sublime Text 3 Editor

| Files 	| Installation 	| Usage 	|
|:-	|:-	|:-:	|
| [add_date.py](add_date.py) 	| Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/`<br>Add these to `Preference` > `Key Bindings` > `User keymap`<br>[<br>    {"keys": ["ctrl+shift+,"], "command": "add_date" },<br>    {"keys": ["ctrl+shift+."], "command": "add_time" }<br>] 	| `Ctrl` + `Shift` + , <br> `Ctrl` + `Shift` + . 	|
| [python_about.sublime-snippet](python_about.sublime-snippet) 	| Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/` 	| about + `Tab` 	|
| [python_shebang.sublime-snippet](python_shebang.sublime-snippet) 	| Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/` 	| shebang + `Tab` 	|
|Monokai.sublime-color-scheme|Package Install `Package Resource Viewer`<br>`Ctrl+Shift+P >> pvr >> Open Resource >> Color Scheme - Default >>` <br> Add `"highlight": "var(blue)",` in `globals`<br>* [Ref1](https://sublimetext.userecho.com/en/communities/1/topics/4674-make-highlight-matches-easier-to-see) <br>* [Ref2](https://www.sublimetext.com/docs/3/color_schemes.html#global_settings-find)|

## Sublime REPL (Side-by-side Build System)
| Files 	| Installation 	| Usage 	|
|:-	|:-	|:-:	|
||[Origami](https://github.com/SublimeText/Origami#installation)|[Usage](https://github.com/SublimeText/Origami#using-the-command-line)
||[Terminus](https://packagecontrol.io/packages/Terminus)|[Usage](https://packagecontrol.io/packages/Terminus)
|[PythonREPL.sublime-build](PythonREPL.sublime-build)|Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/`|
|[terminus_python.sublime-build](terminus_python.sublime-build)|Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/`|
|[terminus_repl.py](terminus_repl.py)|Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/`|
|[newPython3.sublime-build](newPython3.sublime-build)|Save this in `/[home_directory]/.config/sublime-text-3/Packages/User/`|

## Reference
* [[QT08] Supercharging Terminus builds in Sublime Text](https://www.youtube.com/watch?v=HaH3U53UvcY&ab_channel=OdatNurd-SublimeTextTutorials)
````

## File: Sublime-Text-3/terminus_python.sublime-build
````
{
    "target": "terminus_open",
    "cancel": "terminus_cancel_build",
    "cmd": ["/usr/bin/python3", "-u", "$file"],
    "file_regex": "^[ ]File \"(...?)\", line ([0-9]*)",
    "selector": "source.python",
    "working_dir": "$folder",
    "auto_close": false,
    "timeit": true
}
````

## File: Sublime-Text-3/terminus_repl.py
````python
import sublime
import sublime_plugin


def _terminus_view(window, tagName):
    window = window or sublime.active_window()
    for view in window.views():
        if view.settings().get("terminus_view.tag") == tagName:
            return view

    return None


class SendSelectionToTerminusCommand(sublime_plugin.WindowCommand):
    def run(self, tag=None, visible_only=False):
        view = self.window.active_view()
        if view == None:
            return

        if any(sel.empty() for sel in view.sel()):
            spans = [sublime.Region(0, len(view))]
        else:
            spans = [sel for sel in view.sel() if not sel.empty()]

        for span in spans:
            self.window.run_command("terminus_send_string", {
                "string": view.substr(span),
                "tag": tag,
                "visible_only": visible_only
                })


class SampleREPLListener(sublime_plugin.EventListener):
    def on_query_context(self, view, key, operator, operand, match_all):
        if key == "terminus_tag.exists" or key == "terminus_tag.notexists":
            view = _terminus_view(view.window(), operand)
            return view != None if key == "terminus_tag.exists" else view == None

        return None
````

## File: Windows_Terminal/Microsoft.PowerShell_profile.ps1
````powershell
oh-my-posh --init --shell pwsh --config "$env:POSH_THEMES_PATH/catppuccin_mocha.omp.json" | Invoke-Expression

# My Alias

function sage_shell {
    Start-Process "C:\Users\rey\AppData\Local\SageMath 9.3\runtime\bin\mintty.exe" -ArgumentList "/bin/bash --login -c '/opt/sagemath-9.3/sage -sh'"
}

function v {
    nvim $args
}

function e {
    explorer $args
}

function py3 {
    python $args
}
````

## File: Windows_Terminal/README.md
````markdown
## Windows Terminal Settings

### PowerShell 7
 - Download: https://github.com/powershell/powershell/releases
### Nerd Fonts
- [Hurmit](https://www.nerdfonts.com/font-downloads)
- [FiraCode](https://www.nerdfonts.com/font-downloads)
````

## File: Windows_Terminal/settings.json
````json
{
    "$help": "https://aka.ms/terminal-documentation",
    "$schema": "https://aka.ms/terminal-profiles-schema",
    "initialCols": 80,
    "initialRows": 16,
    "actions": 
    [
        {
            "command": 
            {
                "action": "copy",
                "singleLine": false
            },
            "id": "User.copy.644BA8F2",
            "keys": "ctrl+c"
        },
        {
            "command": "paste",
            "id": "User.paste",
            "keys": "ctrl+v"
        },
        {
            "command": "find",
            "id": "User.find",
            "keys": "ctrl+shift+f"
        },
        {
            "command": 
            {
                "action": "splitPane",
                "split": "auto",
                "splitMode": "duplicate"
            },
            "id": "User.splitPane.A6751878",
            "keys": "alt+shift+d"
        }
    ],
    "copyFormatting": "none",
    "copyOnSelect": false,
    "defaultProfile": "{574e775e-4f2a-5b96-ac1e-a2962a402336}",
    "newTabMenu": 
    [
        {
            "type": "remainingProfiles"
        }
    ],
    "profiles": 
    {
        "defaults": 
        {
            "font": 
            {
                "face": "Hurmit Nerd Font Mono"
            },
            "opacity": 85,
            "startingDirectory": "%USERPROFILE%",
            "useAcrylic": false
        },
        "list": 
        [
            {
                "commandline": "C:\\Program Files\\PowerShell\\7\\pwsh.exe -nologo",
                "guid": "{574e775e-4f2a-5b96-ac1e-a2962a402336}",
                "hidden": false,
                "icon": "C:\\Program Files\\PowerShell\\7\\assets\\Powershell_avatar.ico",
                "name": "PowerShell 7",
                "source": "Windows.Terminal.PowershellCore"
            },
            {
                "commandline": "%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\powershell.exe -nologo",
                "guid": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}",
                "hidden": false,
                "name": "Windows PowerShell"
            },
            {
                "commandline": "%SystemRoot%\\System32\\cmd.exe",
                "guid": "{0caa0dad-35be-5f56-a8ff-afceeeaa6101}",
                "hidden": false,
                "name": "Command Prompt"
            },
            {
                "guid": "{b453ae62-4e3d-5e58-b989-0a998ec441b8}",
                "hidden": false,
                "name": "Azure Cloud Shell",
                "source": "Windows.Terminal.Azure"
            },
            {
                "guid": "{2ece5bfe-50ed-5f3a-ab87-5cd4baafed2b}",
                "hidden": false,
                "name": "Git Bash",
                "source": "Git"
            },
            {
                "guid": "{8925b247-95e8-570d-a14f-e32bf534e8ba}",
                "hidden": false,
                "name": "Developer Command Prompt for VS 2022",
                "source": "Windows.Terminal.VisualStudio"
            },
            {
                "guid": "{89d6868d-88f3-5252-b2b6-ac09cec0f0d0}",
                "hidden": false,
                "name": "Developer PowerShell for VS 2022",
                "source": "Windows.Terminal.VisualStudio"
            },
            {
                "guid": "{303bb5b2-9004-5a5c-8ce0-39f54dd21d41}",
                "hidden": false,
                "name": "Developer Command Prompt for VS 2022 (2)",
                "source": "Windows.Terminal.VisualStudio"
            },
            {
                "guid": "{87186f54-d1a7-57f8-928b-f7fea7ca651e}",
                "hidden": false,
                "name": "Developer PowerShell for VS 2022 (2)",
                "source": "Windows.Terminal.VisualStudio"
            }
        ]
    },
    "schemes": [],
    "themes": [],
    "useAcrylicInTabRow": true
}
````
