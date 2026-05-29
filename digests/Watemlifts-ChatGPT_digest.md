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
.github/ISSUE_TEMPLATE/bug_report.md
.github/ISSUE_TEMPLATE/feature_request.md
.github/workflows/codeql.yml
.github/workflows/pylint.yml
.github/workflows/python-package.yml
.github/workflows/python-publish.yml
.gitignore
.pre-commit-config.yaml
AI_watem_agent.ipynb
config.json.example
CONTRIBUTING.md
LICENSE
logo.png
README.md
requirements.txt
SECURITY.md
setup.cfg
setup.py
src/__init__.py
src/revChatGPT/__init__.py
src/revChatGPT/__main__.py
src/revChatGPT/ChatGPT.py
src/revChatGPT/GPTserver.py
src/revChatGPT/Official.py
wiki/revChatGPT.md
wiki/Setup.md
wiki/Star-history.md
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/ISSUE_TEMPLATE/bug_report.md">
---

name: Bug report
about: Create a report to help us improve
title: "[BUG]"
labels: bug
assignees: ''

---

**Description**
A clear and concise description of what the bug is.

**Steps to Reproduce**
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Output**
In the correct directory, run the following command:
`python3 -m revChatGPT --debug`

**Environment**
Please update your packages before reporting the issue:
`pip3 install --upgrade revChatGPT`

 - OS: [e.g. Linux, MacOS, Windows]
 - Python version: python -V
 - ChatGPT Version: pip3 show revChatGPT

**Additional context**
Add any other context about the problem here.
</file>

<file path=".github/ISSUE_TEMPLATE/feature_request.md">
---

name: Feature request
about: Suggest an idea for this project
title: '[Descriptive title]'
labels: ''
assignees: ''

---

**Is your feature request related to a problem? Please describe.**
A clear and concise description of what the problem is.

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Benefits or Justification**
A clear and concise explanation of the benefits of this feature and how it addresses a problem or need.

**Alternatives considered**
A clear and concise description of any alternative solutions or features you've considered, and an explanation of why the proposed solution is preferable.

**Implementation details**
Details about how the feature might be implemented, including any technical challenges, dependencies, integrations, and potential impacts on existing functionality.

**Additional context**
Add any other relevant context or screenshots, mockups, or prototypes here.
</file>

<file path=".github/workflows/codeql.yml">
name: "Code Scanning - Action"

on:
  release:
    types: [published]
  schedule:
    #        ┌───────────── minute (0 - 59)
    #        │  ┌───────────── hour (0 - 23)
    #        │  │ ┌───────────── day of the month (1 - 31)
    #        │  │ │ ┌───────────── month (1 - 12 or JAN-DEC)
    #        │  │ │ │ ┌───────────── day of the week (0 - 6 or SUN-SAT)
    #        │  │ │ │ │
    #        │  │ │ │ │
    #        │  │ │ │ │
    #        *  * * * *
    - cron: '30 1 * * 0'

jobs:
  CodeQL-Build:
    runs-on: ubuntu-latest

    permissions:
      # required for all workflows
      security-events: write
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      # Initializes the CodeQL tools for scanning.
      - name: Initialize CodeQL
        uses: github/codeql-action/init@v2
        with:
          languages: python

      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@v2
</file>

<file path=".github/workflows/pylint.yml">
name: Pylint

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        python-version: ["3.8", "3.9", "3.10"]
    steps:
    - uses: actions/checkout@v3
    - name: Set up Python ${{ matrix.python-version }}
      uses: actions/setup-python@v4
      with:
        python-version: ${{ matrix.python-version }}
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install pylint revChatGPT
    - name: Analysing the code with pylint
      run: |
        pylint $(git ls-files '*.py')
</file>

<file path=".github/workflows/python-package.yml">
# This workflow will install Python dependencies, run tests and lint with a variety of Python versions
# For more information see: https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-python

name: Python package

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  build:

    runs-on: ubuntu-latest
    strategy:
      fail-fast: false
      matrix:
        python-version: ["3.8", "3.9", "3.10"]

    steps:
    - uses: actions/checkout@v3
    - name: Set up Python ${{ matrix.python-version }}
      uses: actions/setup-python@v3
      with:
        python-version: ${{ matrix.python-version }}
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        python -m pip install flake8 pytest
        python -m pip install .
    - name: Lint with flake8
      run: |
        # stop the build if there are Python syntax errors or undefined names
        flake8 . --count --select=E9,F63,F7,F82 --show-source --statistics
        # exit-zero treats all errors as warnings. The GitHub editor is 127 chars wide
        flake8 . --count --exit-zero --max-complexity=10 --max-line-length=127 --statistics
    - name: Test with pytest
      run: pytest
      env:
        OPENAI_EMAIL: ${{ secrets.OPENAI_EMAIL }}
        OPENAI_PASSWORD: ${{ secrets.OPENAI_PASSWORD }}
</file>

<file path=".github/workflows/python-publish.yml">
# This workflow will upload a Python Package using Twine when a release is created
# For more information see: https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-python#publishing-to-package-registries

# This workflow uses actions that are not certified by GitHub.
# They are provided by a third-party and are governed by
# separate terms of service, privacy policy, and support
# documentation.

name: Upload Python Package

on:
  release:
    types: [published]

permissions:
  contents: read

jobs:
  deploy:

    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3
    - name: Set up Python
      uses: actions/setup-python@v4
      with:
        python-version: '3.x'
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install build
    - name: Build package
      run: python -m build
    - name: Publish package
      uses: pypa/gh-action-pypi-publish@27b31702a0e7fc50959f5ad993c78deac1bdfc29
      with:
        user: acheong08
        password: ${{ secrets.PYPI_API_TOKEN }}
</file>

<file path=".pre-commit-config.yaml">
repos:
  - repo: https://github.com/asottile/reorder_python_imports
    rev: v3.9.0
    hooks:
      - id: reorder-python-imports
        args: [--py37-plus]
  - repo: https://github.com/asottile/add-trailing-comma
    rev: v2.3.0
    hooks:
      - id: add-trailing-comma
        args: [--py36-plus]
  - repo: https://github.com/asottile/pyupgrade
    rev: v3.3.1
    hooks:
      - id: pyupgrade
        args: [--py37-plus]

  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: debug-statements
      - id: double-quote-string-fixer
      - id: name-tests-test
      - id: requirements-txt-fixer
  - repo: https://github.com/psf/black
    rev: 22.10.0
    hooks:
      - id: black
</file>

<file path="AI_watem_agent.ipynb">
{
  "cells": [
    {
      "cell_type": "markdown",
      "metadata": {
        "id": "view-in-github",
        "colab_type": "text"
      },
      "source": [
        "<a href=\"https://colab.research.google.com/github/Watemlifts/ChatGPT/blob/main/AI_watem_agent.ipynb\" target=\"_parent\"><img src=\"https://colab.research.google.com/assets/colab-badge.svg\" alt=\"Open In Colab\"/></a>"
      ]
    },
    {
      "cell_type": "markdown",
      "source": [
        "Add `%load_ext cudf.pandas` before importing pandas to speed up operations using GPU"
      ],
      "metadata": {
        "id": "zsvIsZNWzBF3"
      }
    },
    {
      "cell_type": "code",
      "source": [
        "\n",
        "| Tool                               | Free Tier                          | Purpose                    |\n",
        "| ---------------------------------- | ---------------------------------- | -------------------------- |\n",
        "| **Google Colab** (Python notebook) | 12 hrs/day GPU                     | Runs the agent code        |\n",
        "| **OpenAI GPT-3.5-turbo**           | 200 free tokens/min via OpenRouter | Generates text             |\n",
        "| **Google Apps Script**             | Free w/ Gmail                      | Posts to Blogger           |\n",
        "| **YouTube Data API v3**            | 10 000 units/day                   | Posts to YouTube Community |\n"
      ],
      "metadata": {
        "id": "eBJIiAUtzBGe",
        "colab": {
          "base_uri": "https://localhost:8080/",
          "height": 110
        },
        "outputId": "7caffcf1-ace4-42c5-a4da-ddebe73c092f"
      },
      "execution_count": null,
      "outputs": [
        {
          "output_type": "error",
          "ename": "SyntaxError",
          "evalue": "invalid syntax (ipython-input-1-2659166559.py, line 1)",
          "traceback": [
            "\u001b[0;36m  File \u001b[0;32m\"/tmp/ipython-input-1-2659166559.py\"\u001b[0;36m, line \u001b[0;32m1\u001b[0m\n\u001b[0;31m    | Tool                               | Free Tier                          | Purpose                    |\u001b[0m\n\u001b[0m    ^\u001b[0m\n\u001b[0;31mSyntaxError\u001b[0m\u001b[0;31m:\u001b[0m invalid syntax\n"
          ]
        }
      ]
    },
    {
      "cell_type": "code",
      "execution_count": null,
      "metadata": {
        "id": "lIYdn1woOS1n",
        "colab": {
          "base_uri": "https://localhost:8080/",
          "height": 110
        },
        "outputId": "54967c6a-ea24-4443-f28d-827b1708face"
      },
      "outputs": [
        {
          "output_type": "error",
          "ename": "SyntaxError",
          "evalue": "invalid character '“' (U+201C) (ipython-input-2-2713320262.py, line 3)",
          "traceback": [
            "\u001b[0;36m  File \u001b[0;32m\"/tmp/ipython-input-2-2713320262.py\"\u001b[0;36m, line \u001b[0;32m3\u001b[0m\n\u001b[0;31m    | **1. 30-sec Diagnostics**    | “Use a ₹2 000 IR gun to spot a hot motor bearing before it seizes—today’s Kampala case at Oasis Mall.” |\u001b[0m\n\u001b[0m                                     ^\u001b[0m\n\u001b[0;31mSyntaxError\u001b[0m\u001b[0;31m:\u001b[0m invalid character '“' (U+201C)\n"
          ]
        }
      ],
      "source": [
        "| Pillar                       | Example Daily Angle                                                                                    |\n",
        "| ---------------------------- | ------------------------------------------------------------------------------------------------------ |\n",
        "| **1. 30-sec Diagnostics**    | “Use a ₹2 000 IR gun to spot a hot motor bearing before it seizes—today’s Kampala case at Oasis Mall.” |\n",
        "| **2. Energy Hack**           | “Program the VVVF drive to 40 Hz during off-peak hours; saved 18 % kWh last week at Acacia Mall.”      |\n",
        "| **3. Spare-Parts Sourcing**  | “Where to buy genuine Otis door locks in Kampala for < \\$80—vendor WhatsApp inside.”                   |\n",
        "| **4. Compliance Mini-Alert** | “UNBS inspection checklist item #7 you keep failing—door clutch release time.”                         |\n",
        "| **5. Case Study Snapshot**   | “How we cut 22 hrs downtime at Garden City in 3 steps (with PDF logbook).”                             |\n",
        "| **6. Myth-Buster**           | “‘Lubricating rails weekly saves energy’—false! Here’s the lab data.”                                  |\n"
      ]
    },
    {
      "cell_type": "code",
      "source": [
        "#@title Daily Elevator Niche Post Generator\n",
        "!pip install openai --quiet\n",
        "import openai, datetime, json, requests, os\n",
        "\n",
        "openai.api_key = \"sk-XXX\"  # Use OpenRouter key (free)\n",
        "\n",
        "pillars = [\"Diagnostics\",\"Energy\",\"Parts\",\"Compliance\",\"Case\",\"Myth\"]\n",
        "today = datetime.date.today().toordinal()\n",
        "pillar = pillars[today % 6]\n",
        "\n",
        "prompt = f\"\"\"\n",
        "You are a Ugandan lift engineer. Write a 120-word micro-lesson for facility managers on how to {pillar}-related reduce downtime or energy. Include a single actionable tip and end with a CTA to comment.\n",
        "Keep tone conversational. No emojis.\n",
        "\"\"\"\n",
        "\n",
        "res = openai.ChatCompletion.create(\n",
        "  model=\"gpt-3.5-turbo\",\n",
        "  messages=[{\"role\":\"user\",\"content\":prompt}]\n",
        ")\n",
        "text = res[\"choices\"][0][\"message\"][\"content\"]\n",
        "\n",
        "# Save to Google Drive\n",
        "with open(\"/content/today_post.txt\",\"w\") as f:\n",
        "  f.write(text)"
      ],
      "metadata": {
        "id": "HgSgVQKb0Ctw"
      },
      "execution_count": null,
      "outputs": []
    },
    {
      "cell_type": "code",
      "source": [
        "AIzaSyAJN2oxXBX0VfjJQ3SuCnT8KQzxcF9t-Jw"
      ],
      "metadata": {
        "id": "N7-XjOe10KYH"
      },
      "execution_count": null,
      "outputs": []
    },
    {
      "cell_type": "code",
      "source": [
        "from google.auth.transport.requests import Request\n",
        "from google_auth_oauthlib.flow import InstalledAppFlow\n",
        "from googleapiclient.discovery import build\n",
        "import pickle\n",
        "\n",
        "SCOPES = ['https://www.googleapis.com/auth/youtube.force-ssl']\n",
        "creds = None\n",
        "if os.path.exists('token.pickle'):\n",
        "  with open('token.pickle','rb') as f: creds = pickle.load(f)\n",
        "if not creds or not creds.valid:\n",
        "  flow = InstalledAppFlow.from_client_secrets_file('credentials.json', SCOPES)\n",
        "  creds = flow.run_console()\n",
        "  with open('token.pickle','wb') as f: pickle.dump(creds,f)\n",
        "\n",
        "youtube = build('youtube','v3',credentials=creds)\n",
        "youtube.community().insert(\n",
        "  part='snippet',\n",
        "  body={'snippet':{'type':'text','text':text}}\n",
        ").execute()"
      ],
      "metadata": {
        "id": "5JO-EKohAR-A",
        "colab": {
          "base_uri": "https://localhost:8080/",
          "height": 218
        },
        "outputId": "8cca83e4-a481-4a90-cbf6-87601d79a5e9"
      },
      "execution_count": null,
      "outputs": [
        {
          "output_type": "error",
          "ename": "NameError",
          "evalue": "name 'os' is not defined",
          "traceback": [
            "\u001b[0;31m---------------------------------------------------------------------------\u001b[0m",
            "\u001b[0;31mNameError\u001b[0m                                 Traceback (most recent call last)",
            "\u001b[0;32m/tmp/ipython-input-3-2303343143.py\u001b[0m in \u001b[0;36m<cell line: 0>\u001b[0;34m()\u001b[0m\n\u001b[1;32m      6\u001b[0m \u001b[0mSCOPES\u001b[0m \u001b[0;34m=\u001b[0m \u001b[0;34m[\u001b[0m\u001b[0;34m'https://www.googleapis.com/auth/youtube.force-ssl'\u001b[0m\u001b[0;34m]\u001b[0m\u001b[0;34m\u001b[0m\u001b[0;34m\u001b[0m\u001b[0m\n\u001b[1;32m      7\u001b[0m \u001b[0mcreds\u001b[0m \u001b[0;34m=\u001b[0m \u001b[0;32mNone\u001b[0m\u001b[0;34m\u001b[0m\u001b[0;34m\u001b[0m\u001b[0m\n\u001b[0;32m----> 8\u001b[0;31m \u001b[0;32mif\u001b[0m \u001b[0mos\u001b[0m\u001b[0;34m.\u001b[0m\u001b[0mpath\u001b[0m\u001b[0;34m.\u001b[0m\u001b[0mexists\u001b[0m\u001b[0;34m(\u001b[0m\u001b[0;34m'token.pickle'\u001b[0m\u001b[0;34m)\u001b[0m\u001b[0;34m:\u001b[0m\u001b[0;34m\u001b[0m\u001b[0;34m\u001b[0m\u001b[0m\n\u001b[0m\u001b[1;32m      9\u001b[0m   \u001b[0;32mwith\u001b[0m \u001b[0mopen\u001b[0m\u001b[0;34m(\u001b[0m\u001b[0;34m'token.pickle'\u001b[0m\u001b[0;34m,\u001b[0m\u001b[0;34m'rb'\u001b[0m\u001b[0;34m)\u001b[0m \u001b[0;32mas\u001b[0m \u001b[0mf\u001b[0m\u001b[0;34m:\u001b[0m \u001b[0mcreds\u001b[0m \u001b[0;34m=\u001b[0m \u001b[0mpickle\u001b[0m\u001b[0;34m.\u001b[0m\u001b[0mload\u001b[0m\u001b[0;34m(\u001b[0m\u001b[0mf\u001b[0m\u001b[0;34m)\u001b[0m\u001b[0;34m\u001b[0m\u001b[0;34m\u001b[0m\u001b[0m\n\u001b[1;32m     10\u001b[0m \u001b[0;32mif\u001b[0m \u001b[0;32mnot\u001b[0m \u001b[0mcreds\u001b[0m \u001b[0;32mor\u001b[0m \u001b[0;32mnot\u001b[0m \u001b[0mcreds\u001b[0m\u001b[0;34m.\u001b[0m\u001b[0mvalid\u001b[0m\u001b[0;34m:\u001b[0m\u001b[0;34m\u001b[0m\u001b[0;34m\u001b[0m\u001b[0m\n",
            "\u001b[0;31mNameError\u001b[0m: name 'os' is not defined"
          ]
        }
      ]
    }
  ],
  "metadata": {
    "colab": {
      "provenance": [],
      "include_colab_link": true
    },
    "kernelspec": {
      "display_name": "Python 3",
      "name": "python3"
    }
  },
  "nbformat": 4,
  "nbformat_minor": 0
}
</file>

<file path="config.json.example">
{
  "session_token": "<token>",
  "proxy":"<proxy>",
  "accept_language": "en-US,en"
}
</file>

<file path="CONTRIBUTING.md">
# Contributing to ChatGPT

We welcome contributions to ChatGPT! Here are some guidelines to help you get started.

## Types of contributions we are looking for

ChatGPT is an open-source project, and we welcome a wide range of contributions. Here are some examples of the types of contributions we are looking for:

- Code patches
- Documentation improvements
- Bug reports and fixes
- Feature requests and suggestions

Please note that ChatGPT is intended to be used as a development library, so contributions should stay within this scope.

## How to submit a contribution

If you would like to contribute to ChatGPT, follow these steps:

1. Fork the ChatGPT repository.
2. Create a new branch in your fork to make your changes.
3. Commit your changes to your new branch.
4. Push your changes to your fork on GitHub.
5. Submit a pull request from your branch to the ChatGPT repository.

We will review your pull request and, if everything looks good, merge it into the main codebase.

## Questions

If you have any questions about contributing to ChatGPT, feel free to open an issue in the ChatGPT repository and ask.

Thank you for considering a contribution to ChatGPT!
</file>

<file path="LICENSE">
GNU GENERAL PUBLIC LICENSE
                       Version 2, June 1991

 Copyright (C) 1989, 1991 Free Software Foundation, Inc.,
 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.

                            Preamble

  The licenses for most software are designed to take away your
freedom to share and change it.  By contrast, the GNU General Public
License is intended to guarantee your freedom to share and change free
software--to make sure the software is free for all its users.  This
General Public License applies to most of the Free Software
Foundation's software and to any other program whose authors commit to
using it.  (Some other Free Software Foundation software is covered by
the GNU Lesser General Public License instead.)  You can apply it to
your programs, too.

  When we speak of free software, we are referring to freedom, not
price.  Our General Public Licenses are designed to make sure that you
have the freedom to distribute copies of free software (and charge for
this service if you wish), that you receive source code or can get it
if you want it, that you can change the software or use pieces of it
in new free programs; and that you know you can do these things.

  To protect your rights, we need to make restrictions that forbid
anyone to deny you these rights or to ask you to surrender the rights.
These restrictions translate to certain responsibilities for you if you
distribute copies of the software, or if you modify it.

  For example, if you distribute copies of such a program, whether
gratis or for a fee, you must give the recipients all the rights that
you have.  You must make sure that they, too, receive or can get the
source code.  And you must show them these terms so they know their
rights.

  We protect your rights with two steps: (1) copyright the software, and
(2) offer you this license which gives you legal permission to copy,
distribute and/or modify the software.

  Also, for each author's protection and ours, we want to make certain
that everyone understands that there is no warranty for this free
software.  If the software is modified by someone else and passed on, we
want its recipients to know that what they have is not the original, so
that any problems introduced by others will not reflect on the original
authors' reputations.

  Finally, any free program is threatened constantly by software
patents.  We wish to avoid the danger that redistributors of a free
program will individually obtain patent licenses, in effect making the
program proprietary.  To prevent this, we have made it clear that any
patent must be licensed for everyone's free use or not licensed at all.

  The precise terms and conditions for copying, distribution and
modification follow.

                    GNU GENERAL PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. This License applies to any program or other work which contains
a notice placed by the copyright holder saying it may be distributed
under the terms of this General Public License.  The "Program", below,
refers to any such program or work, and a "work based on the Program"
means either the Program or any derivative work under copyright law:
that is to say, a work containing the Program or a portion of it,
either verbatim or with modifications and/or translated into another
language.  (Hereinafter, translation is included without limitation in
the term "modification".)  Each licensee is addressed as "you".

Activities other than copying, distribution and modification are not
covered by this License; they are outside its scope.  The act of
running the Program is not restricted, and the output from the Program
is covered only if its contents constitute a work based on the
Program (independent of having been made by running the Program).
Whether that is true depends on what the Program does.

  1. You may copy and distribute verbatim copies of the Program's
source code as you receive it, in any medium, provided that you
conspicuously and appropriately publish on each copy an appropriate
copyright notice and disclaimer of warranty; keep intact all the
notices that refer to this License and to the absence of any warranty;
and give any other recipients of the Program a copy of this License
along with the Program.

You may charge a fee for the physical act of transferring a copy, and
you may at your option offer warranty protection in exchange for a fee.

  2. You may modify your copy or copies of the Program or any portion
of it, thus forming a work based on the Program, and copy and
distribute such modifications or work under the terms of Section 1
above, provided that you also meet all of these conditions:

    a) You must cause the modified files to carry prominent notices
    stating that you changed the files and the date of any change.

    b) You must cause any work that you distribute or publish, that in
    whole or in part contains or is derived from the Program or any
    part thereof, to be licensed as a whole at no charge to all third
    parties under the terms of this License.

    c) If the modified program normally reads commands interactively
    when run, you must cause it, when started running for such
    interactive use in the most ordinary way, to print or display an
    announcement including an appropriate copyright notice and a
    notice that there is no warranty (or else, saying that you provide
    a warranty) and that users may redistribute the program under
    these conditions, and telling the user how to view a copy of this
    License.  (Exception: if the Program itself is interactive but
    does not normally print such an announcement, your work based on
    the Program is not required to print an announcement.)

These requirements apply to the modified work as a whole.  If
identifiable sections of that work are not derived from the Program,
and can be reasonably considered independent and separate works in
themselves, then this License, and its terms, do not apply to those
sections when you distribute them as separate works.  But when you
distribute the same sections as part of a whole which is a work based
on the Program, the distribution of the whole must be on the terms of
this License, whose permissions for other licensees extend to the
entire whole, and thus to each and every part regardless of who wrote it.

Thus, it is not the intent of this section to claim rights or contest
your rights to work written entirely by you; rather, the intent is to
exercise the right to control the distribution of derivative or
collective works based on the Program.

In addition, mere aggregation of another work not based on the Program
with the Program (or with a work based on the Program) on a volume of
a storage or distribution medium does not bring the other work under
the scope of this License.

  3. You may copy and distribute the Program (or a work based on it,
under Section 2) in object code or executable form under the terms of
Sections 1 and 2 above provided that you also do one of the following:

    a) Accompany it with the complete corresponding machine-readable
    source code, which must be distributed under the terms of Sections
    1 and 2 above on a medium customarily used for software interchange; or,

    b) Accompany it with a written offer, valid for at least three
    years, to give any third party, for a charge no more than your
    cost of physically performing source distribution, a complete
    machine-readable copy of the corresponding source code, to be
    distributed under the terms of Sections 1 and 2 above on a medium
    customarily used for software interchange; or,

    c) Accompany it with the information you received as to the offer
    to distribute corresponding source code.  (This alternative is
    allowed only for noncommercial distribution and only if you
    received the program in object code or executable form with such
    an offer, in accord with Subsection b above.)

The source code for a work means the preferred form of the work for
making modifications to it.  For an executable work, complete source
code means all the source code for all modules it contains, plus any
associated interface definition files, plus the scripts used to
control compilation and installation of the executable.  However, as a
special exception, the source code distributed need not include
anything that is normally distributed (in either source or binary
form) with the major components (compiler, kernel, and so on) of the
operating system on which the executable runs, unless that component
itself accompanies the executable.

If distribution of executable or object code is made by offering
access to copy from a designated place, then offering equivalent
access to copy the source code from the same place counts as
distribution of the source code, even though third parties are not
compelled to copy the source along with the object code.

  4. You may not copy, modify, sublicense, or distribute the Program
except as expressly provided under this License.  Any attempt
otherwise to copy, modify, sublicense or distribute the Program is
void, and will automatically terminate your rights under this License.
However, parties who have received copies, or rights, from you under
this License will not have their licenses terminated so long as such
parties remain in full compliance.

  5. You are not required to accept this License, since you have not
signed it.  However, nothing else grants you permission to modify or
distribute the Program or its derivative works.  These actions are
prohibited by law if you do not accept this License.  Therefore, by
modifying or distributing the Program (or any work based on the
Program), you indicate your acceptance of this License to do so, and
all its terms and conditions for copying, distributing or modifying
the Program or works based on it.

  6. Each time you redistribute the Program (or any work based on the
Program), the recipient automatically receives a license from the
original licensor to copy, distribute or modify the Program subject to
these terms and conditions.  You may not impose any further
restrictions on the recipients' exercise of the rights granted herein.
You are not responsible for enforcing compliance by third parties to
this License.

  7. If, as a consequence of a court judgment or allegation of patent
infringement or for any other reason (not limited to patent issues),
conditions are imposed on you (whether by court order, agreement or
otherwise) that contradict the conditions of this License, they do not
excuse you from the conditions of this License.  If you cannot
distribute so as to satisfy simultaneously your obligations under this
License and any other pertinent obligations, then as a consequence you
may not distribute the Program at all.  For example, if a patent
license would not permit royalty-free redistribution of the Program by
all those who receive copies directly or indirectly through you, then
the only way you could satisfy both it and this License would be to
refrain entirely from distribution of the Program.

If any portion of this section is held invalid or unenforceable under
any particular circumstance, the balance of the section is intended to
apply and the section as a whole is intended to apply in other
circumstances.

It is not the purpose of this section to induce you to infringe any
patents or other property right claims or to contest validity of any
such claims; this section has the sole purpose of protecting the
integrity of the free software distribution system, which is
implemented by public license practices.  Many people have made
generous contributions to the wide range of software distributed
through that system in reliance on consistent application of that
system; it is up to the author/donor to decide if he or she is willing
to distribute software through any other system and a licensee cannot
impose that choice.

This section is intended to make thoroughly clear what is believed to
be a consequence of the rest of this License.

  8. If the distribution and/or use of the Program is restricted in
certain countries either by patents or by copyrighted interfaces, the
original copyright holder who places the Program under this License
may add an explicit geographical distribution limitation excluding
those countries, so that distribution is permitted only in or among
countries not thus excluded.  In such case, this License incorporates
the limitation as if written in the body of this License.

  9. The Free Software Foundation may publish revised and/or new versions
of the General Public License from time to time.  Such new versions will
be similar in spirit to the present version, but may differ in detail to
address new problems or concerns.

Each version is given a distinguishing version number.  If the Program
specifies a version number of this License which applies to it and "any
later version", you have the option of following the terms and conditions
either of that version or of any later version published by the Free
Software Foundation.  If the Program does not specify a version number of
this License, you may choose any version ever published by the Free Software
Foundation.

  10. If you wish to incorporate parts of the Program into other free
programs whose distribution conditions are different, write to the author
to ask for permission.  For software which is copyrighted by the Free
Software Foundation, write to the Free Software Foundation; we sometimes
make exceptions for this.  Our decision will be guided by the two goals
of preserving the free status of all derivatives of our free software and
of promoting the sharing and reuse of software generally.

                            NO WARRANTY

  11. BECAUSE THE PROGRAM IS LICENSED FREE OF CHARGE, THERE IS NO WARRANTY
FOR THE PROGRAM, TO THE EXTENT PERMITTED BY APPLICABLE LAW.  EXCEPT WHEN
OTHERWISE STATED IN WRITING THE COPYRIGHT HOLDERS AND/OR OTHER PARTIES
PROVIDE THE PROGRAM "AS IS" WITHOUT WARRANTY OF ANY KIND, EITHER EXPRESSED
OR IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.  THE ENTIRE RISK AS
TO THE QUALITY AND PERFORMANCE OF THE PROGRAM IS WITH YOU.  SHOULD THE
PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF ALL NECESSARY SERVICING,
REPAIR OR CORRECTION.

  12. IN NO EVENT UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING
WILL ANY COPYRIGHT HOLDER, OR ANY OTHER PARTY WHO MAY MODIFY AND/OR
REDISTRIBUTE THE PROGRAM AS PERMITTED ABOVE, BE LIABLE TO YOU FOR DAMAGES,
INCLUDING ANY GENERAL, SPECIAL, INCIDENTAL OR CONSEQUENTIAL DAMAGES ARISING
OUT OF THE USE OR INABILITY TO USE THE PROGRAM (INCLUDING BUT NOT LIMITED
TO LOSS OF DATA OR DATA BEING RENDERED INACCURATE OR LOSSES SUSTAINED BY
YOU OR THIRD PARTIES OR A FAILURE OF THE PROGRAM TO OPERATE WITH ANY OTHER
PROGRAMS), EVEN IF SUCH HOLDER OR OTHER PARTY HAS BEEN ADVISED OF THE
POSSIBILITY OF SUCH DAMAGES.

                     END OF TERMS AND CONDITIONS

            How to Apply These Terms to Your New Programs

  If you develop a new program, and you want it to be of the greatest
possible use to the public, the best way to achieve this is to make it
free software which everyone can redistribute and change under these terms.

  To do so, attach the following notices to the program.  It is safest
to attach them to the start of each source file to most effectively
convey the exclusion of warranty; and each file should have at least
the "copyright" line and a pointer to where the full notice is found.

    <one line to give the program's name and a brief idea of what it does.>
    Copyright (C) <year>  <name of author>

    This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation; either version 2 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License along
    with this program; if not, write to the Free Software Foundation, Inc.,
    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

Also add information on how to contact you by electronic and paper mail.

If the program is interactive, make it output a short notice like this
when it starts in an interactive mode:

    Gnomovision version 69, Copyright (C) year name of author
    Gnomovision comes with ABSOLUTELY NO WARRANTY; for details type `show w'.
    This is free software, and you are welcome to redistribute it
    under certain conditions; type `show c' for details.

The hypothetical commands `show w' and `show c' should show the appropriate
parts of the General Public License.  Of course, the commands you use may
be called something other than `show w' and `show c'; they could even be
mouse-clicks or menu items--whatever suits your program.

You should also get your employer (if you work as a programmer) or your
school, if any, to sign a "copyright disclaimer" for the program, if
necessary.  Here is a sample; alter the names:

  Yoyodyne, Inc., hereby disclaims all copyright interest in the program
  `Gnomovision' (which makes passes at compilers) written by James Hacker.

  <signature of Ty Coon>, 1 April 1989
  Ty Coon, President of Vice

This General Public License does not permit incorporating your program into
proprietary programs.  If your program is a subroutine library, you may
consider it more useful to permit linking proprietary applications with the
library.  If this is what you want to do, use the GNU Lesser General
Public License instead of this License.
</file>

<file path="requirements.txt">
openai
pre-commit
tiktoken
</file>

<file path="SECURITY.md">
# Security Policy

## Supported Versions

Use this section to tell people about which versions of your project are
currently being supported with security updates.

| Version | Supported          |
| ------- | ------------------ |
| 5.1.x   | :white_check_mark: |
| 5.0.x   | :x:                |
| 4.0.x   | :white_check_mark: |
| < 4.0   | :x:                |

## Reporting a Vulnerability

Use this section to tell people how to report a vulnerability.

Tell them where to go, how often they can expect to get an update on a
reported vulnerability, what to expect if the vulnerability is accepted or
declined, etc.
</file>

<file path="setup.cfg">
[metadata]
description_file=README.md
license_files=LICENSE
</file>

<file path="setup.py">
from setuptools import find_packages
from setuptools import setup

setup(
    name="revChatGPT",
    version="1.1.3",
    license="GNU General Public License v2.0",
    author="Antonio Cheong",
    author_email="acheong@student.dalat.org",
    description="ChatGPT is a reverse engineering of OpenAI's ChatGPT API",
    packages=find_packages("src"),
    package_dir={"": "src"},
    py_modules=["revChatGPT", "GPTserver", "Official"],
    url="https://github.com/acheong08/ChatGPT",
    install_requires=[
        "openai",
        "tiktoken",
    ],
    # optional dependencies
    extras_require={
        "api": ["flask"],
        "unofficial": [
            "undetected_chromedriver>=3.1.7",
            "selenium>=4.7.2",
            "tls_client>=0.1.7",
            "2captcha-python>=1.1.3",
        ],
    },
    long_description=open("README.md", encoding="utf-8").read(),
    long_description_content_type="text/markdown",
    entry_points={
        "console_scripts": [
            "revChatGPT = revChatGPT.__main__:main",
            "revApiGPT = revChatGPT.GPTserver:main",
            "OfficialChatGPT = revChatGPT.Official:main",
        ],
    },
)
</file>

<file path="src/__init__.py">

</file>

<file path="src/revChatGPT/__init__.py">

</file>

<file path="src/revChatGPT/__main__.py">
import json
from os import getenv
from os.path import exists

from revChatGPT.ChatGPT import Chatbot


def get_input(prompt):
    # Display the prompt
    print(prompt, end="")

    # Initialize an empty list to store the input lines
    lines = []

    # Read lines of input until the user enters an empty line
    while True:
        line = input()
        if line == "":
            break
        lines.append(line)

    # Join the lines, separated by newlines, and store the result
    user_input = "\n".join(lines)

    # Return the input
    return user_input


def configure():
    config_files = ["config.json"]
    xdg_config_home = getenv("XDG_CONFIG_HOME")
    if xdg_config_home:
        config_files.append(f"{xdg_config_home}/revChatGPT/config.json")
    user_home = getenv("HOME")
    if user_home:
        config_files.append(f"{user_home}/.config/revChatGPT/config.json")

    config_file = next((f for f in config_files if exists(f)), None)
    if config_file:
        with open(config_file, encoding="utf-8") as f:
            config = json.load(f)
    else:
        print("No config file found.")
        raise Exception("No config file found.")
    return config


def chatGPT_main(config):
    print("Logging in...")
    chatbot = Chatbot(config)
    while True:
        prompt = get_input("\nYou:\n")
        if prompt.startswith("!"):
            if prompt == "!help":
                print(
                    """
                !help - Show this message
                !reset - Forget the current conversation
                !refresh - Refresh the session authentication
                !config - Show the current configuration
                !rollback x - Rollback the conversation (x being the number of messages to rollback)
                !exit - Exit this program
                """,
                )
                continue
            elif prompt == "!reset":
                chatbot.reset_chat()
                print("Chat session successfully reset.")
                continue
            elif prompt == "!refresh":
                chatbot.refresh_session()
                print("Session successfully refreshed.\n")
                continue
            elif prompt == "!config":
                print(json.dumps(chatbot.config, indent=4))
                continue
            elif prompt.startswith("!rollback"):
                # Default to 1 rollback if no number is specified
                try:
                    rollback = int(prompt.split(" ")[1])
                except IndexError:
                    rollback = 1
                chatbot.rollback_conversation(rollback)
                print(f"Rolled back {rollback} messages.")
                continue
            elif prompt.startswith("!setconversation"):
                try:
                    chatbot.config["conversation"] = prompt.split(" ")[1]
                    print("Conversation has been changed")
                except IndexError:
                    print("Please include conversation UUID in command")
                continue
            elif prompt == "!exit":
                break
        try:
            print("Chatbot: ")
            message = chatbot.ask(
                prompt,
                conversation_id=chatbot.config.get("conversation"),
                parent_id=chatbot.config.get("parent_id"),
            )
            print(message["message"])
        except Exception as exc:
            print("Something went wrong!")
            print(exc)
            continue


def main():
    print(
        """
        ChatGPT - A command-line interface to OpenAI's ChatGPT (https://chat.openai.com/chat)
        Repo: github.com/acheong08/ChatGPT
        """,
    )
    print("Type '!help' to show a full list of commands")
    print("Press enter twice to submit your question.\n")
    chatGPT_main(configure())


if __name__ == "__main__":
    main()
</file>

<file path="src/revChatGPT/ChatGPT.py">
import json
import logging
import re
import uuid
from time import sleep

import tls_client
import undetected_chromedriver as uc
from requests.exceptions import HTTPError
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait
from twocaptcha import TwoCaptcha

# Disable all logging
logging.basicConfig(level=logging.ERROR)

BASE_URL = "https://chat.openai.com/"


class Chrome(uc.Chrome):
    def __del__(self):
        self.quit()


class Chatbot:
    def __init__(
        self,
        config,
        conversation_id=None,
        parent_id=None,
        no_refresh=False,
    ) -> None:
        self.config = config
        self.session = tls_client.Session(
            client_identifier="chrome_108",
        )
        if "proxy" in config:
            if type(config["proxy"]) != str:
                raise Exception("Proxy must be a string!")
            proxies = {
                "http": config["proxy"],
                "https": config["proxy"],
            }
            self.session.proxies.update(proxies)
        if "verbose" in config:
            if type(config["verbose"]) != bool:
                raise Exception("Verbose must be a boolean!")
            self.verbose = config["verbose"]
        else:
            self.verbose = False
        self.conversation_id = conversation_id
        self.parent_id = parent_id
        self.conversation_mapping = {}
        self.conversation_id_prev_queue = []
        self.parent_id_prev_queue = []
        self.isMicrosoftLogin = False
        self.twocaptcha_key = None
        # stdout colors
        self.GREEN = "\033[92m"
        self.WARNING = "\033[93m"
        self.ENDCOLOR = "\033[0m"
        if "email" in config and "password" in config:
            if type(config["email"]) != str:
                raise Exception("Email must be a string!")
            if type(config["password"]) != str:
                raise Exception("Password must be a string!")
            self.email = config["email"]
            self.password = config["password"]
            if "isMicrosoftLogin" in config and config["isMicrosoftLogin"] == True:
                self.isMicrosoftLogin = True
                self.microsoft_login()
            elif "captcha" in config:
                if type(config["captcha"]) != str:
                    raise Exception("2Captcha API Key must be a string!")
                self.twocaptcha_key = config["captcha"]
                self.email_login(self.solve_captcha())
            else:
                raise Exception("Invalid config!")
        elif "session_token" in config:
            if no_refresh:
                self.get_cf_cookies()
                return
            if type(config["session_token"]) != str:
                raise Exception("Session token must be a string!")
            self.session_token = config["session_token"]
            self.session.cookies.set(
                "__Secure-next-auth.session-token",
                config["session_token"],
            )
            self.get_cf_cookies()
        else:
            raise Exception("Invalid config!")
        self.retry_refresh()

    def retry_refresh(self):
        retries = 5
        refresh = True
        while refresh:
            try:
                self.refresh_session()
                refresh = False
            except Exception as exc:
                if retries == 0:
                    raise exc
                retries -= 1

    def ask(
        self,
        prompt,
        conversation_id=None,
        parent_id=None,
        gen_title=False,
        session_token=None,
    ):
        if session_token:
            self.session.cookies.set(
                "__Secure-next-auth.session-token",
                session_token,
            )
            self.session_token = session_token
            self.config["session_token"] = session_token
        self.retry_refresh()
        self.map_conversations()
        if conversation_id == None:
            conversation_id = self.conversation_id
        if parent_id == None:
            parent_id = (
                self.parent_id
                if conversation_id == self.conversation_id
                else self.conversation_mapping[conversation_id]
            )
        data = {
            "action": "next",
            "messages": [
                {
                    "id": str(uuid.uuid4()),
                    "role": "user",
                    "content": {"content_type": "text", "parts": [prompt]},
                },
            ],
            "conversation_id": conversation_id,
            "parent_message_id": parent_id or str(uuid.uuid4()),
            "model": "text-davinci-002-render",
        }
        new_conv = data["conversation_id"] is None
        self.conversation_id_prev_queue.append(
            data["conversation_id"],
        )  # for rollback
        self.parent_id_prev_queue.append(data["parent_message_id"])
        response = self.session.post(
            url=BASE_URL + "backend-api/conversation",
            data=json.dumps(data),
            timeout_seconds=180,
        )
        if response.status_code != 200:
            print(response.text)
            self.refresh_session()
            raise HTTPError(
                f"Wrong response code: {response.status_code}! Refreshing session...",
            )
        else:
            try:
                response = response.text.splitlines()[-4]
                response = response[6:]
            except Exception as exc:
                print("Incorrect response from OpenAI API")
                raise Exception("Incorrect response from OpenAI API") from exc
            # Check if it is JSON
            if response.startswith("{"):
                response = json.loads(response)
                self.parent_id = response["message"]["id"]
                self.conversation_id = response["conversation_id"]
                message = response["message"]["content"]["parts"][0]
                res = {
                    "message": message,
                    "conversation_id": self.conversation_id,
                    "parent_id": self.parent_id,
                }
                if gen_title and new_conv:
                    try:
                        title = self.gen_title(
                            self.conversation_id,
                            self.parent_id,
                        )["title"]
                    except Exception as exc:
                        split = prompt.split(" ")
                        title = " ".join(split[:3]) + ("..." if len(split) > 3 else "")
                    res["title"] = title
                return res
            else:
                return None

    def check_response(self, response):
        if response.status_code != 200:
            print(response.text)
            raise Exception("Response code error: ", response.status_code)

    def get_conversations(self, offset=0, limit=20):
        url = BASE_URL + f"backend-api/conversations?offset={offset}&limit={limit}"
        response = self.session.get(url)
        self.check_response(response)
        data = json.loads(response.text)
        return data["items"]

    def get_msg_history(self, id):
        url = BASE_URL + f"backend-api/conversation/{id}"
        response = self.session.get(url)
        self.check_response(response)
        data = json.loads(response.text)
        return data

    def gen_title(self, id, message_id):
        url = BASE_URL + f"backend-api/conversation/gen_title/{id}"
        response = self.session.post(
            url,
            data=json.dumps(
                {"message_id": message_id, "model": "text-davinci-002-render"},
            ),
        )
        self.check_response(response)
        data = json.loads(response.text)
        return data

    def change_title(self, id, title):
        url = BASE_URL + f"backend-api/conversation/{id}"
        response = self.session.patch(url, data=f'{{"title": "{title}"}}')
        self.check_response(response)

    def delete_conversation(self, id):
        url = BASE_URL + f"backend-api/conversation/{id}"
        response = self.session.patch(url, data='{"is_visible": false}')
        self.check_response(response)

    def clear_conversations(self):
        url = BASE_URL + "backend-api/conversations"
        response = self.session.patch(url, data='{"is_visible": false}')
        self.check_response(response)

    def map_conversations(self):
        conversations = self.get_conversations()
        histories = [self.get_msg_history(x["id"]) for x in conversations]
        for x, y in zip(conversations, histories):
            self.conversation_mapping[x["id"]] = y["current_node"]

    def refresh_session(self, session_token=None):
        if session_token:
            self.session.cookies.set(
                "__Secure-next-auth.session-token",
                session_token,
            )
            self.session_token = session_token
            self.config["session_token"] = session_token
        url = BASE_URL + "api/auth/session"
        response = self.session.get(url, timeout_seconds=180)
        if response.status_code == 403:
            self.get_cf_cookies()
            raise Exception("Clearance refreshing...")
        try:
            if "error" in response.json():
                raise Exception(
                    f"Failed to refresh session! Error: {response.json()['error']}",
                )
            elif (
                response.status_code != 200
                or response.json() == {}
                or "accessToken" not in response.json()
            ):
                raise Exception(
                    f"Response code: {response.status_code} \n Response: {response.text}",
                )
            else:
                self.session.headers.update(
                    {
                        "Authorization": "Bearer " + response.json()["accessToken"],
                    },
                )
            self.session_token = self.session.cookies._find(
                "__Secure-next-auth.session-token",
            )
        except Exception as exc:
            print("Failed to refresh session!")
            if self.isMicrosoftLogin:
                print("Attempting to re-authenticate...")
                self.microsoft_login()
            elif self.twocaptcha_key:
                self.email_login(self.solve_captcha())
            else:
                raise Exception("Failed to refresh session!") from exc

    def reset_chat(self) -> None:
        """
        Reset the conversation ID and parent ID.

        :return: None
        """
        self.conversation_id = None
        self.parent_id = str(uuid.uuid4())

    def microsoft_login(self) -> None:
        """
        Login to OpenAI via Microsoft Login Authentication.

        :return: None
        """
        try:
            # Open the browser
            self.cf_cookie_found = False
            self.session_cookie_found = False
            self.agent_found = False
            self.cf_clearance = None
            self.user_agent = None
            options = self.__get_ChromeOptions()
            print("Spawning browser...")
            driver = uc.Chrome(
                enable_cdp_events=True,
                options=options,
                driver_executable_path=self.config.get("driver_exec_path"),
                browser_executable_path=self.config.get("browser_exec_path"),
            )
            print("Browser spawned.")
            driver.add_cdp_listener(
                "Network.responseReceivedExtraInfo",
                lambda msg: self.detect_cookies(msg),
            )
            driver.add_cdp_listener(
                "Network.requestWillBeSentExtraInfo",
                lambda msg: self.detect_user_agent(msg),
            )
            driver.get(BASE_URL)
            while not self.agent_found or not self.cf_cookie_found:
                sleep(5)
            self.refresh_headers(
                cf_clearance=self.cf_clearance,
                user_agent=self.user_agent,
            )
            # Wait for the login button to appear
            WebDriverWait(driver, 120).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//button[contains(text(), 'Log in')]"),
                ),
            )
            # Click the login button
            driver.find_element(
                by=By.XPATH,
                value="//button[contains(text(), 'Log in')]",
            ).click()
            # Wait for the Login with Microsoft button to be clickable
            WebDriverWait(driver, 60).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//button[@data-provider='windowslive']"),
                ),
            )
            # Click the Login with Microsoft button
            driver.find_element(
                by=By.XPATH,
                value="//button[@data-provider='windowslive']",
            ).click()
            # Wait for the email input field to appear
            WebDriverWait(driver, 60).until(
                EC.visibility_of_element_located(
                    (By.XPATH, "//input[@type='email']"),
                ),
            )
            # Enter the email
            driver.find_element(
                by=By.XPATH,
                value="//input[@type='email']",
            ).send_keys(self.config["email"])
            # Wait for the Next button to be clickable
            WebDriverWait(driver, 60).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//input[@type='submit']"),
                ),
            )
            # Click the Next button
            driver.find_element(
                by=By.XPATH,
                value="//input[@type='submit']",
            ).click()
            # Wait for the password input field to appear
            WebDriverWait(driver, 60).until(
                EC.visibility_of_element_located(
                    (By.XPATH, "//input[@type='password']"),
                ),
            )
            # Enter the password
            driver.find_element(
                by=By.XPATH,
                value="//input[@type='password']",
            ).send_keys(self.config["password"])
            # Wait for the Sign in button to be clickable
            WebDriverWait(driver, 60).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//input[@type='submit']"),
                ),
            )
            # Click the Sign in button
            driver.find_element(
                by=By.XPATH,
                value="//input[@type='submit']",
            ).click()
            # Wait for the Allow button to appear
            WebDriverWait(driver, 60).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//input[@type='submit']"),
                ),
            )
            # click Yes button
            driver.find_element(
                by=By.XPATH,
                value="//input[@type='submit']",
            ).click()
            # wait for input box to appear (to make sure we're signed in)
            WebDriverWait(driver, 60).until(
                EC.visibility_of_element_located(
                    (By.XPATH, "//textarea"),
                ),
            )
            while not self.session_cookie_found:
                sleep(5)
            print(self.GREEN + "Login successful." + self.ENDCOLOR)
        finally:
            # Close the browser
            driver.quit()
            del driver

    def solve_captcha(self) -> str:
        """
        Solve the 2Captcha captcha.

        :return: str
        """
        twocaptcha_key = self.twocaptcha_key
        twocaptcha_solver_config = {
            "apiKey": twocaptcha_key,
            "defaultTimeout": 120,
            "recaptchaTimeout": 600,
            "pollingInterval": 10,
        }
        twocaptcha_solver = TwoCaptcha(**twocaptcha_solver_config)
        print("Waiting for captcha to be solved...")
        solved_captcha = twocaptcha_solver.recaptcha(
            sitekey="6Lc-wnQjAAAAADa5SPd68d0O3xmj0030uaVzpnXP",
            url="https://auth0.openai.com/u/login/identifier",
        )
        if "code" in solved_captcha:
            print(self.GREEN + "Captcha solved successfully!" + self.ENDCOLOR)
            if self.verbose:
                print(
                    self.GREEN
                    + "Captcha token: "
                    + self.ENDCOLOR
                    + solved_captcha["code"],
                )
            return solved_captcha

    def email_login(self, solved_captcha) -> None:
        """
        Login to OpenAI via Email/Password Authentication and 2Captcha.

        :return: None
        """
        # Open the browser
        try:
            self.cf_cookie_found = False
            self.session_cookie_found = False
            self.agent_found = False
            self.cf_clearance = None
            self.user_agent = None
            options = self.__get_ChromeOptions()
            print("Spawning browser...")
            driver = uc.Chrome(
                enable_cdp_events=True,
                options=options,
                driver_executable_path=self.config.get("driver_exec_path"),
                browser_executable_path=self.config.get("browser_exec_path"),
            )
            print("Browser spawned.")
            driver.add_cdp_listener(
                "Network.responseReceivedExtraInfo",
                lambda msg: self.detect_cookies(msg),
            )
            driver.add_cdp_listener(
                "Network.requestWillBeSentExtraInfo",
                lambda msg: self.detect_user_agent(msg),
            )
            driver.get(BASE_URL)
            while not self.agent_found or not self.cf_cookie_found:
                sleep(5)
            self.refresh_headers(
                cf_clearance=self.cf_clearance,
                user_agent=self.user_agent,
            )
            # Wait for the login button to appear
            WebDriverWait(driver, 120).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//button[contains(text(), 'Log in')]"),
                ),
            )
            # Click the login button
            driver.find_element(
                by=By.XPATH,
                value="//button[contains(text(), 'Log in')]",
            ).click()
            # Wait for the email input field to appear
            WebDriverWait(driver, 60).until(
                EC.visibility_of_element_located(
                    (By.ID, "username"),
                ),
            )
            # Enter the email
            driver.find_element(by=By.ID, value="username").send_keys(
                self.config["email"],
            )
            # Wait for Recaptcha to appear
            WebDriverWait(driver, 60).until(
                EC.presence_of_element_located(
                    (By.CSS_SELECTOR, "*[name*='g-recaptcha-response']"),
                ),
            )
            # Find Recaptcha
            google_captcha_response_input = driver.find_element(
                By.CSS_SELECTOR,
                "*[name*='g-recaptcha-response']",
            )
            captcha_input = driver.find_element(By.NAME, "captcha")
            # Make input visible
            driver.execute_script(
                "arguments[0].setAttribute('style','type: text; visibility:visible;');",
                google_captcha_response_input,
            )
            driver.execute_script(
                "arguments[0].setAttribute('style','type: text; visibility:visible;');",
                captcha_input,
            )
            driver.execute_script(
                """
            document.getElementById("g-recaptcha-response").innerHTML = arguments[0]
            """,
                solved_captcha.get("code"),
            )
            driver.execute_script(
                """
            document.querySelector("input[name='captcha']").value = arguments[0]
            """,
                solved_captcha.get("code"),
            )
            # Hide the captcha input
            driver.execute_script(
                "arguments[0].setAttribute('style', 'display:none;');",
                google_captcha_response_input,
            )
            # Wait for the Continue button to be clickable
            WebDriverWait(driver, 60).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//button[@type='submit']"),
                ),
            )
            # Click the Continue button
            driver.find_element(
                by=By.XPATH,
                value="//button[@type='submit']",
            ).click()
            # Wait for the password input field to appear
            WebDriverWait(driver, 60).until(
                EC.visibility_of_element_located(
                    (By.ID, "password"),
                ),
            )
            # Enter the password
            driver.find_element(by=By.ID, value="password").send_keys(
                self.config["password"],
            )
            # Wait for the Sign in button to be clickable
            WebDriverWait(driver, 60).until(
                EC.element_to_be_clickable(
                    (By.XPATH, "//button[@type='submit']"),
                ),
            )
            # Click the Sign in button
            driver.find_element(
                by=By.XPATH,
                value="//button[@type='submit']",
            ).click()
            # wait for input box to appear (to make sure we're signed in)
            WebDriverWait(driver, 60).until(
                EC.visibility_of_element_located(
                    (By.XPATH, "//textarea"),
                ),
            )
            while not self.session_cookie_found:
                sleep(5)
            print(self.GREEN + "Login successful." + self.ENDCOLOR)
        finally:
            # Close the browser
            driver.quit()
            del driver

    def __get_ChromeOptions(self):
        options = uc.ChromeOptions()
        options.add_argument("--start_maximized")
        options.add_argument("--disable-extensions")
        options.add_argument("--disable-application-cache")
        options.add_argument("--disable-gpu")
        options.add_argument("--no-sandbox")
        options.add_argument("--disable-setuid-sandbox")
        options.add_argument("--disable-dev-shm-usage")
        if self.config.get("proxy", "") != "":
            options.add_argument("--proxy-server=" + self.config["proxy"])
        return options

    def get_cf_cookies(self) -> None:
        """
        Get cloudflare cookies.

        :return: None
        """
        try:
            self.cf_cookie_found = False
            self.agent_found = False
            self.cf_clearance = None
            self.user_agent = None
            options = self.__get_ChromeOptions()
            print("Spawning browser...")
            driver = uc.Chrome(
                enable_cdp_events=True,
                options=options,
                driver_executable_path=self.config.get("driver_exec_path"),
                browser_executable_path=self.config.get("browser_exec_path"),
            )
            print("Browser spawned.")
            driver.add_cdp_listener(
                "Network.responseReceivedExtraInfo",
                lambda msg: self.detect_cookies(msg),
            )
            driver.add_cdp_listener(
                "Network.requestWillBeSentExtraInfo",
                lambda msg: self.detect_user_agent(msg),
            )
            driver.get("https://chat.openai.com/chat")
            while not self.agent_found or not self.cf_cookie_found:
                sleep(5)
        finally:
            # Close the browser
            driver.quit()
            del driver
            self.refresh_headers(
                cf_clearance=self.cf_clearance,
                user_agent=self.user_agent,
            )

    def detect_cookies(self, message):
        if "params" in message:
            if "headers" in message["params"]:
                if "set-cookie" in message["params"]["headers"]:
                    # Use regex to get the cookie for cf_clearance=*;
                    cf_clearance_cookie = re.search(
                        "cf_clearance=.*?;",
                        message["params"]["headers"]["set-cookie"],
                    )
                    session_cookie = re.search(
                        "__Secure-next-auth.session-token=.*?;",
                        message["params"]["headers"]["set-cookie"],
                    )
                    if cf_clearance_cookie and not self.cf_cookie_found:
                        print("Found Cloudflare Cookie!")
                        # remove the semicolon and 'cf_clearance=' from the string
                        raw_cf_cookie = cf_clearance_cookie.group(0)
                        self.cf_clearance = raw_cf_cookie.split("=")[1][:-1]
                        if self.verbose:
                            print(
                                self.GREEN
                                + "Cloudflare Cookie: "
                                + self.ENDCOLOR
                                + self.cf_clearance,
                            )
                        self.cf_cookie_found = True
                    if session_cookie and not self.session_cookie_found:
                        print("Found Session Token!")
                        # remove the semicolon and '__Secure-next-auth.session-token=' from the string
                        raw_session_cookie = session_cookie.group(0)
                        self.session_token = raw_session_cookie.split("=")[1][:-1]
                        self.session.cookies.set(
                            "__Secure-next-auth.session-token",
                            self.session_token,
                        )
                        if self.verbose:
                            print(
                                self.GREEN
                                + "Session Token: "
                                + self.ENDCOLOR
                                + self.session_token,
                            )
                        self.session_cookie_found = True

    def detect_user_agent(self, message):
        if "params" in message:
            if "headers" in message["params"]:
                if "user-agent" in message["params"]["headers"]:
                    # Use regex to get the cookie for cf_clearance=*;
                    user_agent = message["params"]["headers"]["user-agent"]
                    self.user_agent = user_agent
                    self.agent_found = True
        self.refresh_headers(
            cf_clearance=self.cf_clearance,
            user_agent=self.user_agent,
        )

    def refresh_headers(self, cf_clearance, user_agent):
        del self.session.cookies["cf_clearance"]
        self.session.headers.clear()
        self.session.cookies.set("cf_clearance", cf_clearance)
        self.session.headers.update(
            {
                "Accept": "text/event-stream",
                "Authorization": "Bearer ",
                "Content-Type": "application/json",
                "User-Agent": user_agent,
                "X-Openai-Assistant-App-Id": "",
                "Connection": "close",
                "Accept-Language": "en-US,en;q=0.9",
                "Referer": "https://chat.openai.com/chat",
            },
        )

    def rollback_conversation(self, num=1) -> None:
        """
        Rollback the conversation.
        :param num: The number of messages to rollback
        :return: None
        """
        for i in range(num):
            self.conversation_id = self.conversation_id_prev_queue.pop()
            self.parent_id = self.parent_id_prev_queue.pop()
</file>

<file path="src/revChatGPT/GPTserver.py">
from flask import Flask
from flask import jsonify
from flask import request
from revChatGPT.ChatGPT import Chatbot

app = Flask(__name__)

# Session token based rate limiting
token_available: dict = {}

chatbot = Chatbot(config={}, conversation_id=None, parent_id=None, no_refresh=True)


def verify_data(data: dict) -> bool:
    """
    Verifies that the required fields are present in the data.
    """
    # Required fields: "prompt", "session_token"
    if "prompt" not in data or "session_token" not in data:
        return False
    return True


@app.route("/chat", methods=["POST"])
def chat():
    """
    The main chat endpoint.
    """
    data: dict = request.get_json()
    if not verify_data(data=data):
        return jsonify({"error": "Invalid data."}), 400

    chatbot.session_token: str = data["session_token"]

    conversation_id = data.get("conversation_id", None)
    parent_id = data.get("parent_id", None)

    # Return rate limit if token_available is false
    if token_available.get(
        data.get("session_token"),
    ) != None and not token_available.get(data.get("session_token")):
        return jsonify({"error": "Rate limited"}), 429

    token_available[data.get("session_token")] = False

    try:
        response = chatbot.ask(
            prompt=data["prompt"],
            session_token=data["session_token"],
            parent_id=parent_id,
            conversation_id=conversation_id,
        )
    except Exception as exc:
        token_available[data.get("session_token")] = True
        return jsonify({"error": str(exc)}), 500

    response["session_token"] = chatbot.session_token

    token_available[data.get("session_token")] = True

    return jsonify(response), 200


@app.route("/refresh", methods=["POST"])
def refresh():
    """
    The refresh endpoint.
    """
    data = request.get_json()
    if "session_token" not in data:
        return jsonify({"error": "Invalid data."}), 400
    if not token_available.get(data.get("session_token")):
        return jsonify({"error": "Invalid token."}), 400
    chatbot.session_token = data["session_token"]
    try:
        chatbot.refresh_session()
    except Exception as exc:
        return jsonify({"error": str(exc)}), 400

    return jsonify({"session_token": chatbot.session_token}), 200


def main():
    app.run(host="127.0.0.1", port=8080)
</file>

<file path="src/revChatGPT/Official.py">
"""
A simple wrapper for the official ChatGPT API
"""
import argparse
import json
import os
import sys
from datetime import date

import openai
import tiktoken

ENGINE = os.environ.get("GPT_ENGINE") or "text-chat-davinci-002-20221122"

ENCODER = tiktoken.get_encoding("gpt2")


def get_max_tokens(prompt: str) -> int:
    """
    Get the max tokens for a prompt
    """
    return 4000 - len(ENCODER.encode(prompt))


class Chatbot:
    """
    Official ChatGPT API
    """

    def __init__(self, api_key: str, buffer: int = None) -> None:
        """
        Initialize Chatbot with API key (from https://platform.openai.com/account/api-keys)
        """
        openai.api_key = api_key or os.environ.get("OPENAI_API_KEY")
        self.conversations = Conversation()
        self.prompt = Prompt(buffer=buffer)

    def _get_completion(
        self,
        prompt: str,
        temperature: float = 0.5,
        stream: bool = False,
    ):
        """
        Get the completion function
        """
        return openai.Completion.create(
            engine=ENGINE,
            prompt=prompt,
            temperature=temperature,
            max_tokens=get_max_tokens(prompt),
            stop=["\n\n\n"],
            stream=stream,
        )

    def _process_completion(
        self, user_request: str, completion: dict, conversation_id: str = None, user: str = "User"
    ) -> dict:
        if completion.get("choices") is None:
            raise Exception("ChatGPT API returned no choices")
        if len(completion["choices"]) == 0:
            raise Exception("ChatGPT API returned no choices")
        if completion["choices"][0].get("text") is None:
            raise Exception("ChatGPT API returned no text")
        completion["choices"][0]["text"] = completion["choices"][0]["text"].rstrip(
            "<|im_end|>",
        )
        # Add to chat history
        self.prompt.add_to_history(user_request, completion["choices"][0]["text"], user=user)
        if conversation_id is not None:
            self.save_conversation(conversation_id)
        return completion

    def _process_completion_stream(
        self, user_request: str, completion: dict, conversation_id: str = None, user: str = "User"
    ) -> str:
        full_response = ""
        for response in completion:
            if response.get("choices") is None:
                raise Exception("ChatGPT API returned no choices")
            if len(response["choices"]) == 0:
                raise Exception("ChatGPT API returned no choices")
            if response["choices"][0].get("finish_details") is not None:
                break
            if response["choices"][0].get("text") is None:
                raise Exception("ChatGPT API returned no text")
            if response["choices"][0]["text"] == "<|im_end|>":
                break
            yield response["choices"][0]["text"]
            full_response += response["choices"][0]["text"]

        # Add to chat history
        self.prompt.add_to_history(user_request, full_response, user)
        if conversation_id is not None:
            self.save_conversation(conversation_id)

    def ask(
        self, user_request: str, temperature: float = 0.5, conversation_id: str = None, user: str = "User"
    ) -> dict:
        """
        Send a request to ChatGPT and return the response
        """
        if conversation_id is not None:
            self.load_conversation(conversation_id)
        completion = self._get_completion(
            self.prompt.construct_prompt(user_request, user=user),
            temperature,
        )
        return self._process_completion(user_request, completion, user=user)

    def ask_stream(
        self, user_request: str, temperature: float = 0.5, conversation_id: str = None, user: str = "User"
    ) -> str:
        """
        Send a request to ChatGPT and yield the response
        """
        if conversation_id is not None:
            self.load_conversation(conversation_id)
        prompt = self.prompt.construct_prompt(user_request, user=user)
        return self._process_completion_stream(
            user_request=user_request,
            completion=self._get_completion(prompt, temperature, stream=True),
            user=user,
        )

    def make_conversation(self, conversation_id: str) -> None:
        """
        Make a conversation
        """
        self.conversations.add_conversation(conversation_id, [])

    def rollback(self, num: int) -> None:
        """
        Rollback chat history num times
        """
        for _ in range(num):
            self.prompt.chat_history.pop()

    def reset(self) -> None:
        """
        Reset chat history
        """
        self.prompt.chat_history = []

    def load_conversation(self, conversation_id) -> None:
        """
        Load a conversation from the conversation history
        """
        if conversation_id not in self.conversations.conversations:
            # Create a new conversation
            self.make_conversation(conversation_id)
        self.prompt.chat_history = self.conversations.get_conversation(conversation_id)

    def save_conversation(self, conversation_id) -> None:
        """
        Save a conversation to the conversation history
        """
        self.conversations.add_conversation(conversation_id, self.prompt.chat_history)


class AsyncChatbot(Chatbot):
    """
    Official ChatGPT API (async)
    """

    async def _get_completion(
        self,
        prompt: str,
        temperature: float = 0.5,
        stream: bool = False,
    ):
        """
        Get the completion function
        """
        return openai.Completion.acreate(
            engine=ENGINE,
            prompt=prompt,
            temperature=temperature,
            max_tokens=get_max_tokens(prompt),
            stop=["\n\n\n"],
            stream=stream,
        )

    async def ask(self, user_request: str, temperature: float = 0.5, user: str = "User") -> dict:
        """
        Same as Chatbot.ask but async
        }
        """
        completion = await self._get_completion(
            self.prompt.construct_prompt(user_request, user=user),
            temperature,
        )
        return self._process_completion(user_request, completion, user=user)

    async def ask_stream(self, user_request: str, temperature: float = 0.5, user: str = "User") -> str:
        """
        Same as Chatbot.ask_stream but async
        """
        prompt = self.prompt.construct_prompt(user_request, user=user)
        return self._process_completion_stream(
            user_request=user_request,
            completion=await self._get_completion(prompt, temperature, stream=True),
            user=user,
        )


class Prompt:
    """
    Prompt class with methods to construct prompt
    """

    def __init__(self, buffer: int = None) -> None:
        """
        Initialize prompt with base prompt
        """
        self.base_prompt = (
            os.environ.get("CUSTOM_BASE_PROMPT")
            or "You are ChatGPT, a large language model trained by OpenAI. Respond conversationally. Do not answer as the user. Current date: "
            + str(date.today())
            + "\n\n"
            + "User: Hello\n"
            + "ChatGPT: Hello! How can I help you today? <|im_end|>\n\n\n"
        )
        # Track chat history
        self.chat_history: list = []
        self.buffer = buffer

    def add_to_chat_history(self, chat: str) -> None:
        """
        Add chat to chat history for next prompt
        """
        self.chat_history.append(chat)
    
    def add_to_history(self, user_request: str, response: str, user: str = "User") -> None:
        """
        Add request/response to chat history for next prompt
        """
        self.add_to_chat_history(
            user+": "
            + user_request
            + "\n\n\n"
            + "ChatGPT: "
            + response
            + "<|im_end|>\n"
        )

    def history(self, custom_history: list = None) -> str:
        """
        Return chat history
        """
        return "\n".join(custom_history or self.chat_history)

    def construct_prompt(self, new_prompt: str, custom_history: list = None, user: str = "User") -> str:
        """
        Construct prompt based on chat history and request
        """
        prompt = (
            self.base_prompt
            + self.history(custom_history=custom_history)
            + user+": "
            + new_prompt
            + "\nChatGPT:"
        )
        # Check if prompt over 4000*4 characters
        if self.buffer is not None:
            max_tokens = 4000 - self.buffer
        else:
            max_tokens = 3200
        if len(ENCODER.encode(prompt)) > max_tokens:
            # Remove oldest chat
            self.chat_history.pop(0)
            # Construct prompt again
            prompt = self.construct_prompt(new_prompt, custom_history, user)
        return prompt


class Conversation:
    """
    For handling multiple conversations
    """

    def __init__(self) -> None:
        self.conversations = {}

    def add_conversation(self, key: str, history: list) -> None:
        """
        Adds a history list to the conversations dict with the id as the key
        """
        self.conversations[key] = history

    def get_conversation(self, key: str) -> list:
        """
        Retrieves the history list from the conversations dict with the id as the key
        """
        return self.conversations[key]

    def remove_conversation(self, key: str) -> None:
        """
        Removes the history list from the conversations dict with the id as the key
        """
        del self.conversations[key]

    def __str__(self) -> str:
        """
        Creates a JSON string of the conversations
        """
        return json.dumps(self.conversations)

    def save(self, file: str) -> None:
        """
        Saves the conversations to a JSON file
        """
        with open(file, "w", encoding="utf-8") as f:
            f.write(str(self))

    def load(self, file: str) -> None:
        """
        Loads the conversations from a JSON file
        """
        with open(file, encoding="utf-8") as f:
            self.conversations = json.loads(f.read())


def main():
    print(
        """
    ChatGPT - A command-line interface to OpenAI's ChatGPT (https://chat.openai.com/chat)
    Repo: github.com/acheong08/ChatGPT
    """,
    )
    print("Type '!help' to show a full list of commands")
    print("Press enter twice to submit your question.\n")

    def get_input(prompt):
        """
        Multi-line input function
        """
        # Display the prompt
        print(prompt, end="")

        # Initialize an empty list to store the input lines
        lines = []

        # Read lines of input until the user enters an empty line
        while True:
            line = input()
            if line == "":
                break
            lines.append(line)

        # Join the lines, separated by newlines, and store the result
        user_input = "\n".join(lines)

        # Return the input
        return user_input

    def chatbot_commands(cmd: str) -> bool:
        """
        Handle chatbot commands
        """
        if cmd == "!help":
            print(
                """
            !help - Display this message
            !rollback - Rollback chat history
            !reset - Reset chat history
            !prompt - Show current prompt
            !save_c <conversation_name> - Save history to a conversation
            !load_c <conversation_name> - Load history from a conversation
            !save_f <file_name> - Save all conversations to a file
            !load_f <file_name> - Load all conversations from a file
            !exit - Quit chat
            """,
            )
        elif cmd == "!exit":
            exit()
        elif cmd == "!rollback":
            chatbot.rollback(1)
        elif cmd == "!reset":
            chatbot.reset()
        elif cmd == "!prompt":
            print(chatbot.prompt.construct_prompt(""))
        elif cmd.startswith("!save_c"):
            chatbot.save_conversation(cmd.split(" ")[1])
        elif cmd.startswith("!load_c"):
            chatbot.load_conversation(cmd.split(" ")[1])
        elif cmd.startswith("!save_f"):
            chatbot.conversations.save(cmd.split(" ")[1])
        elif cmd.startswith("!load_f"):
            chatbot.conversations.load(cmd.split(" ")[1])
        else:
            return False
        return True

    # Get API key from command line
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--api_key",
        type=str,
        required=True,
        help="OpenAI API key",
    )
    parser.add_argument(
        "--stream",
        action="store_true",
        help="Stream response",
    )
    parser.add_argument(
        "--temperature",
        type=float,
        default=0.5,
        help="Temperature for response",
    )
    args = parser.parse_args()
    # Initialize chatbot
    chatbot = Chatbot(api_key=args.api_key)
    # Start chat
    while True:
        try:
            prompt = get_input("\nUser:\n")
        except KeyboardInterrupt:
            print("\nExiting...")
            sys.exit()
        if prompt.startswith("!"):
            if chatbot_commands(prompt):
                continue
        if not args.stream:
            response = chatbot.ask(prompt, temperature=args.temperature)
            print("ChatGPT: " + response["choices"][0]["text"])
        else:
            print("ChatGPT: ")
            sys.stdout.flush()
            for response in chatbot.ask_stream(prompt, temperature=args.temperature):
                print(response, end="")
                sys.stdout.flush()
            print()


if __name__ == "__main__":
    main()
</file>

<file path="wiki/revChatGPT.md">
<a id="revChatGPT.Official"></a>

# revChatGPT.Official

A simple wrapper for the official ChatGPT API

<a id="revChatGPT.Official.get_max_tokens"></a>

#### get\_max\_tokens

```python
def get_max_tokens(prompt: str) -> int
```

Get the max tokens for a prompt

<a id="revChatGPT.Official.Chatbot"></a>

## Chatbot Objects

```python
class Chatbot()
```

Official ChatGPT API

<a id="revChatGPT.Official.Chatbot.__init__"></a>

#### \_\_init\_\_

```python
def __init__(api_key: str, buffer: int = None) -> None
```

Initialize Chatbot with API key (from https://platform.openai.com/account/api-keys)

<a id="revChatGPT.Official.Chatbot.ask"></a>

#### ask

```python
def ask(user_request: str, temperature: float = 0.5) -> dict
```

Send a request to ChatGPT and return the response

<a id="revChatGPT.Official.Chatbot.ask_stream"></a>

#### ask\_stream

```python
def ask_stream(user_request: str, temperature: float = 0.5) -> str
```

Send a request to ChatGPT and yield the response

<a id="revChatGPT.Official.Chatbot.rollback"></a>

#### rollback

```python
def rollback(num: int) -> None
```

Rollback chat history num times

<a id="revChatGPT.Official.Chatbot.reset"></a>

#### reset

```python
def reset() -> None
```

Reset chat history

<a id="revChatGPT.Official.Chatbot.load_conversation"></a>

#### load\_conversation

```python
def load_conversation(conversation_id) -> None
```

Load a conversation from the conversation history

<a id="revChatGPT.Official.Chatbot.save_conversation"></a>

#### save\_conversation

```python
def save_conversation(conversation_id) -> None
```

Save a conversation to the conversation history

<a id="revChatGPT.Official.AsyncChatbot"></a>

## AsyncChatbot Objects

```python
class AsyncChatbot(Chatbot)
```

Official ChatGPT API (async)

<a id="revChatGPT.Official.AsyncChatbot.ask"></a>

#### ask

```python
async def ask(user_request: str, temperature: float = 0.5) -> dict
```

Same as Chatbot.ask but async
}

<a id="revChatGPT.Official.AsyncChatbot.ask_stream"></a>

#### ask\_stream

```python
async def ask_stream(user_request: str, temperature: float = 0.5) -> str
```

Same as Chatbot.ask_stream but async

<a id="revChatGPT.Official.Prompt"></a>

## Prompt Objects

```python
class Prompt()
```

Prompt class with methods to construct prompt

<a id="revChatGPT.Official.Prompt.__init__"></a>

#### \_\_init\_\_

```python
def __init__(buffer: int = None) -> None
```

Initialize prompt with base prompt

<a id="revChatGPT.Official.Prompt.add_to_chat_history"></a>

#### add\_to\_chat\_history

```python
def add_to_chat_history(chat: str) -> None
```

Add chat to chat history for next prompt

<a id="revChatGPT.Official.Prompt.history"></a>

#### history

```python
def history() -> str
```

Return chat history

<a id="revChatGPT.Official.Prompt.construct_prompt"></a>

#### construct\_prompt

```python
def construct_prompt(new_prompt: str) -> str
```

Construct prompt based on chat history and request

<a id="revChatGPT.Official.Conversation"></a>

## Conversation Objects

```python
class Conversation()
```

For handling multiple conversations

<a id="revChatGPT.Official.Conversation.add_conversation"></a>

#### add\_conversation

```python
def add_conversation(key: str, history: list) -> None
```

Adds a history list to the conversations dict with the id as the key

<a id="revChatGPT.Official.Conversation.get_conversation"></a>

#### get\_conversation

```python
def get_conversation(key: str) -> list
```

Retrieves the history list from the conversations dict with the id as the key

<a id="revChatGPT.Official.Conversation.remove_conversation"></a>

#### remove\_conversation

```python
def remove_conversation(key: str) -> None
```

Removes the history list from the conversations dict with the id as the key

<a id="revChatGPT.Official.Conversation.__str__"></a>

#### \_\_str\_\_

```python
def __str__() -> str
```

Creates a JSON string of the conversations

<a id="revChatGPT.Official.Conversation.save"></a>

#### save

```python
def save(file: str) -> None
```

Saves the conversations to a JSON file

<a id="revChatGPT.Official.Conversation.load"></a>

#### load

```python
def load(file: str) -> None
```

Loads the conversations from a JSON file
</file>

<file path="wiki/Setup.md">
# Setup

`pip3 install --upgrade revChatGPT`

(MacOS might need `brew install --cask chromedriver`. View [#380](https://github.com/acheong08/ChatGPT/issues/380))

## Dependencies

Make sure Chrome or Chromium is installed

If you need to select a different version of chrome/chromium, use the `driver_exec_path` and `browser_exec_path` config options

## Authentication:

You must define the session token or (email and password) for Microsoft Login in the config:

- ### Session Token Authentication:

  You can find the session token manually from your browser:

  1. Go to `https://chat.openai.com/api/auth/session`
  2. Press `F12` to open console
  3. Go to `Application` > `Cookies`
  4. Copy the session token value in `__Secure-next-auth.session-token`
  5. Paste it into `config.json` in the current working directory

  ```json
  { "session_token": "<YOUR_TOKEN>" }
  ```

- ### Email/Password Login Authentication:

  ```json
  {
    "email": "<EMAIL>",
    "password": "<PASSWORD>",
    "captcha": "<2CAPTCHA_API_KEY>"
  }
  ```

  **Note: 2Captcha is required for Email/Password Login**

- ### Microsoft Login Authentication:

  ```json
  {
    "email": "<EMAIL>",
    "password": "<PASSWORD>",
    "isMicrosoftLogin": True
  }
  ```
```
True → Python dict
true → JSON
```

  **Note: `email` and `password` parameters will override `session_token`**

## Server Config:

You can use `Xvfb` to emulate a a display buffer.

https://github.com/ultrafunkamsterdam/undetected-chromedriver/issues/743#issuecomment-1366847803

# Config options (Optional)

```json
{
  "session_token": "<token>",
  "email": "<EMAIL>",
  "password": "<PASSWORD>",
  "captcha": "<2CAPTCHA_API_KEY>",
  "isMicrosoftLogin": True | False,
  "proxy": "<proxy>",
  "driver_exec_path": "./path/to/driver",
  "browser_exec_path": "./path/to/browser",
  "conversation": "<DEFAULT CONVERSATION UUID>",
  "parent_id": "<DEFAULT PARENT ID>"
  "verbose": True | False
}
```
It is impossible to easily get the parent_id and conversation_id from the website. You can only get it programmatically. Don't mess with it unless you know what you are doing

`"driver_exec_path": "/usr/local/bin/chromedriver"` might be necessary for MacOS
```
True → Python dict
true → JSON
```
</file>

<file path="wiki/Star-history.md">
This is just for me to keep track so I can show off to my friends (If I had any...)

[![Star History Chart](https://api.star-history.com/svg?repos=acheong08/ChatGPT&type=Date)](https://star-history.com/#acheong08/ChatGPT&Date)
</file>

<file path=".gitignore">
# Custom
config.json
# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
pip-wheel-metadata/
share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# PyInstaller
#  Usually these files are written by a python script from a template
#  before PyInstaller builds the exe, so as to inject date/other infos into it.
*.manifest
*.spec

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
htmlcov/
.tox/
.nox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.py,cover
.hypothesis/
.pytest_cache/

# Translations
*.mo
*.pot

# Django stuff:
*.log
local_settings.py
db.sqlite3
db.sqlite3-journal

# Flask stuff:
instance/
.webassets-cache

# Scrapy stuff:
.scrapy

# Sphinx documentation
docs/_build/

# PyBuilder
target/

# Jupyter Notebook
.ipynb_checkpoints

# IPython
profile_default/
ipython_config.py

# pyenv
.python-version

# pipenv
#   According to pypa/pipenv#598, it is recommended to include Pipfile.lock in version control.
#   However, in case of collaboration, if having platform-specific dependencies or dependencies
#   having no cross-platform support, pipenv may install dependencies that don't work, or not
#   install all needed dependencies.
#Pipfile.lock

# PEP 582; used by e.g. github.com/David-OConnor/pyflow
__pypackages__/

# Celery stuff
celerybeat-schedule
celerybeat.pid

# SageMath parsed files
*.sage.py

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# Spyder project settings
.spyderproject
.spyproject

# Rope project settings
.ropeproject

# mkdocs documentation
/site

# mypy
.mypy_cache/
.dmypy.json
dmypy.json

# Pyre type checker
.pyre/

# JetBrains IDEs configuration
.idea/
</file>

<file path="README.md">
# ChatGPT <img src="https://github.com/acheong08/ChatGPT/blob/main/logo.png?raw=true" width="7%"></img>

[![PyPi](https://img.shields.io/pypi/v/revChatGPT.svg)](https://pypi.python.org/pypi/revChatGPT)
[![Downloads](https://static.pepy.tech/badge/revchatgpt)](https://pypi.python.org/pypi/revChatGPT)

Reverse Engineered ChatGPT API by OpenAI. Extensible for chatbots etc.

Connect with me on [Linkedin](https://www.linkedin.com/in/acheong08/) to support this project. (Not open for commercial opportunities yet. Too busy)
<br><br>
You can also follow me on [Twitter](https://twitter.com/GodlyIgnorance) to stay up to date.

<details>
<summary>

# Official API (Browserless)

COMPLETELY FREE AND NO RATE LIMITS (Unpatched Bug - Might be fixed later)

</summary>

## Installation
`pip3 install revChatGPT`

## Setup

1. Create account on [OpenAI](https://platform.openai.com/)
2. Go to https://platform.openai.com/account/api-keys
3. Copy API key

## Usage

### Command line
`OfficialChatGPT --api_key API_KEY --stream` (Assumes Python PyPi in PATH)

<details>
<summary>

### Developer
</summary>

Both Async and Sync are available. You can also stream responses via a generator. Read example code to learn more

#### Example code

You can find it [here](https://github.com/acheong08/ChatGPT/blob/main/src/revChatGPT/Official.py#L292-L408)

#### Further Documentation
In [wiki](https://github.com/acheong08/ChatGPT/wiki/revChatGPT)

#### Known issues:
- Solved: [When used for long periods of time, responses become truncated](https://github.com/acheong08/ChatGPT/issues/519)

</details>
</details>

<details>
<summary>

# Reversed API (Browser required)

This breaks terms of service

</summary>

## Installation
`pip3 install revChatGPT[unofficial]`

## Configuration

Refer to the setup [guide](https://github.com/acheong08/ChatGPT/wiki/Setup) for more information.

## Usage

### Command line

`python3 -m revChatGPT`

```
!help - Show this message
!reset - Forget the current conversation
!refresh - Refresh the session authentication
!config - Show the current configuration
!rollback x - Rollback the conversation (x being the number of messages to rollback)
!exit - Exit this program
```

### API
`python3 -m GPTserver`

HTTP POST request:

```json
{
  "session_token": "eyJhbGciOiJkaXIiL...",
  "prompt": "Your prompt here"
}
```

Optional:

```json
{
  "session_token": "eyJhbGciOiJkaXIiL...",
  "prompt": "Your prompt here",
  "conversation_id": "UUID...",
  "parent_id": "UUID..."
}
```

- Rate limiting is enabled by default to prevent simultaneous requests

### Developer

```python
from revChatGPT.ChatGPT import Chatbot

chatbot = Chatbot({
  "session_token": "<YOUR_TOKEN>"
}, conversation_id=None, parent_id=None) # You can start a custom conversation

response = chatbot.ask("Prompt", conversation_id=None, parent_id=None) # You can specify custom conversation and parent ids. Otherwise it uses the saved conversation (yes. conversations are automatically saved)

print(response)
# {
#   "message": message,
#   "conversation_id": self.conversation_id,
#   "parent_id": self.parent_id,
# }
```

</details>

# Q&A

Q: Is it the real ChatGPT or just a GPT-3 based ripoff?

A: It is the real ChatGPT model found though an info leak on chat.openai.com (patched)

Q: Where did you get the prompt for ChatGPT?

A: https://www.reddit.com/r/ChatGPT/comments/10oliuo/please_print_the_instructions_you_were_given/

Q: <Open pull request with question and I will answer them here -- if significant enough>

# Awesome ChatGPT

[My list](https://github.com/stars/acheong08/lists/awesome-chatgpt)

If you have a cool project you want added to the list, open an issue.

# Disclaimers

This is not an official OpenAI product. This is a personal project and is not affiliated with OpenAI in any way. Don't sue me

# Credits

- [virtualharby](https://twitter.com/virtualharby) - Memes for emotional support
- [rawandahmad698](https://github.com/rawandahmad698) - Reverse engineering Auth0
- [FlorianREGAZ](https://github.com/FlorianREGAZ) - TLS client
- [PyRo1121](https://github.com/PyRo1121) - Linting
- [Harry-Jing](https://github.com/Harry-Jing) - Async support
- [Ukenn2112](https://github.com/Ukenn2112) - Documentation
- [aliferouss19](https://github.com/aliferouss19) - Logo
- [All other contributors](https://github.com/acheong08/ChatGPT/graphs/contributors)
</file>

</files>
