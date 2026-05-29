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
.dockerignore
.env.example
.github/workflows/bettergpt-docker-publish.yml
.github/workflows/docker-publish.yml
.gitignore
docker-compose/bettergpt/docker-compose.yaml
docker-compose/bettergpt/Dockerfile
docker-compose/chatgpt-next-web/docker-compose.yaml
docker-compose/lobe-chat/docker-compose.yaml
docker-compose/README.md
Dockerfile
LICENSE
package.json
README.md
src/app.ts
start.bat
start.sh
tsconfig.json
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".dockerignore">
docker-compose/
.github/
.gitignore
start.bat
Dockerfile
.dockerignore
</file>

<file path=".env.example">
SERVER_PORT=3040
USER_AGENT="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"
NEW_SESSION_RETRIES=5
API_KEY=
CLOUDFLARED=true

PROXY=false
PROXY_HOST=proxy.example.com
PROXY_PORT=8080
PROXY_AUTH=false
PROXY_USERNAME=
PROXY_PASSWORD=
PROXY_PROTOCOL=http
</file>

<file path=".github/workflows/bettergpt-docker-publish.yml">
name: ci

on:
  schedule:
    - cron:  '0 0 * * *' # Runs at 12:00 AM every day

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          file: ./docker-compose/bettergpt/Dockerfile
          push: true
          platforms: linux/amd64,linux/arm64
          tags: pawanosman/bettergpt:latest
</file>

<file path=".github/workflows/docker-publish.yml">
name: ci

on:
  push:
    branches:
      - "main"

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          file: ./Dockerfile
          push: true
          platforms: linux/amd64,linux/arm64
          tags: pawanosman/chatgpt:latest
</file>

<file path="docker-compose/bettergpt/docker-compose.yaml">
version: '3.8' # Specify Docker Compose file version

services:
  # Define the service for the web interface of ChatGPT
  bettergpt:
    image: pawanosman/bettergpt:latest # Use the specified Docker image for the web interface
    ports:
      - "5173:5173" # Map port 5173 on the host to port 5173 in the container
    environment: # Set environment variables for the container
      VITE_CUSTOM_API_ENDPOINT: "http://localhost:3040/v1/chat/completions"
      VITE_DEFAULT_API_ENDPOINT: "http://localhost:3040/v1/chat/completions"
      VITE_OPENAI_API_KEY: "anything"
    depends_on:
      - chatgpt # Ensure this service starts after the chatgpt service

  # Define the backend service for ChatGPT
  chatgpt:
    image: pawanosman/chatgpt:latest # Use the specified Docker image for the backend
    restart: always # Ensure the container restarts automatically if it stops
    ports:
      - "3040:3040" # Map port 3040 on the host to port 3040 in the container
</file>

<file path="docker-compose/bettergpt/Dockerfile">
# Use a Node.js base image
FROM node:19-alpine

# Set the working directory
WORKDIR /app

# Install git
RUN apk add --no-cache git

# Clone the project repository
RUN git clone https://github.com/ztjhz/BetterChatGPT.git /app

# Install dependencies
RUN npm install

# Expose the port the app runs on
EXPOSE 5173

# Command to run the start script
CMD ["npm", "run", "dev"]
</file>

<file path="docker-compose/chatgpt-next-web/docker-compose.yaml">
version: '3.8' # Specify Docker Compose file version

services:
  # Define the service for the web interface of ChatGPT
  chatgpt-next-web:
    image: yidadaa/chatgpt-next-web # Use the specified Docker image for the web interface
    ports:
      - "3000:3000" # Map port 3000 on the host to port 3000 in the container
    environment: # Set environment variables for the container
      OPENAI_API_KEY: "anything" # Placeholder for the actual OpenAI API key
      BASE_URL: "http://chatgpt:3040" # URL for the backend service
      CUSTOM_MODELS: "-all,+gpt-3.5-turbo" # Enable only the gpt-3.5-turbo model, disable all others
    depends_on:
      - chatgpt # Ensure this service starts after the chatgpt service

  # Define the backend service for ChatGPT
  chatgpt:
    image: pawanosman/chatgpt:latest # Use the specified Docker image for the backend
    restart: always # Ensure the container restarts automatically if it stops
    ports:
      - "3040:3040" # Map port 3040 on the host to port 3040 in the container
</file>

<file path="docker-compose/lobe-chat/docker-compose.yaml">
version: '3.8' # Specify Docker Compose file version

services:
  # Define the service for the web interface of ChatGPT
  lobe-chat:
    image: lobehub/lobe-chat:latest # Use the specified Docker image for the web interface
    restart: always
    ports:
      - "3210:3210" # Map port 3210 on the host to port 3210 in the container
    environment: # Set environment variables for the container
      OPENAI_API_KEY: "anything" # Placeholder for the actual OpenAI API key
      OPENAI_PROXY_URL: "http://chatgpt:3040/v1" # URL for the backend service
    depends_on:
      - chatgpt # Ensure this service starts after the chatgpt service

  # Define the backend service for ChatGPT
  chatgpt:
    image: pawanosman/chatgpt:latest # Use the specified Docker image for the backend
    restart: always # Ensure the container restarts automatically if it stops
    ports:
      - "3040:3040" # Map port 3040 on the host to port 3040 in the container
</file>

<file path="docker-compose/README.md">
# Docker Compose Deployment Guide

This guide provides instructions on how to deploy the Docker Compose applications within this project.

## Prerequisites

Before proceeding, ensure you have Docker and Docker Compose installed on your system:

- Docker: [Get Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Deployment Instructions

Within this project, there are multiple Docker Compose applications. Each can be deployed using its respective `docker-compose.yaml` file. Here's how to deploy each one:

Certainly! You can use `curl` or `wget` to download the Docker Compose file directly from a given URL (if the files are hosted online and have a direct URL). Here's how you can do it in a one-liner command for each service, followed by the command to run it:

###  [BetterChatGPT](https://github.com/ztjhz/BetterChatGPT)

```sh
curl -L -o docker-compose.yaml https://raw.githubusercontent.com/PawanOsman/ChatGPT/main/docker-compose/bettergpt/docker-compose.yaml
docker-compose -f docker-compose.yaml up -d
```

Or if you're using `wget`:

```sh
wget https://raw.githubusercontent.com/PawanOsman/ChatGPT/main/docker-compose/bettergpt/docker-compose.yaml -O docker-compose.yaml
docker-compose -f docker-compose.yaml up -d
```

### [ChatGPT Next Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web)

```sh
curl -L -o docker-compose.yaml https://raw.githubusercontent.com/PawanOsman/ChatGPT/main/docker-compose/chatgpt-next-web/docker-compose.yaml
docker-compose -f docker-compose.yaml up -d
```

Or with `wget`:

```sh
wget https://raw.githubusercontent.com/PawanOsman/ChatGPT/main/docker-compose/chatgpt-next-web/docker-compose.yaml -O docker-compose.yaml
docker-compose -f docker-compose.yaml up -d
```

### [Lobe Chat](https://github.com/lobehub/lobe-chat)

```sh
curl -L -o docker-compose.yaml https://raw.githubusercontent.com/PawanOsman/ChatGPT/main/docker-compose/lobe-chat/docker-compose.yaml
docker-compose -f docker-compose.yaml up -d
```

Or using `wget`:

```sh
wget https://raw.githubusercontent.com/PawanOsman/ChatGPT/main/docker-compose/lobe-chat/docker-compose.yaml -O docker-compose.yaml
docker-compose -f docker-compose.yaml up -d
```

## Managing the Applications

Once deployed, you can manage your applications with the following commands:

- To view the status of your services:
  ```sh
  docker-compose ps
  ```

- To stop the services:
  ```sh
  docker-compose down
  ```

- To view the logs of a service:
  ```sh
  docker-compose logs [service-name]
  ```

Replace `[service-name]` with the name of the service you want to check the logs for.

## Additional Notes

- Ensure you are in the correct directory before running the `docker-compose` commands.
- Use the `-d` flag to run containers in detached mode.
- To pull the latest images before starting containers, use the command `docker-compose pull`.

Thank you for using this project. Please report any issues or provide feedback to the project maintainers.
</file>

<file path="Dockerfile">
# Use a Node.js base image
FROM node:19-alpine

# Set the working directory
WORKDIR /app

# Copy source code
COPY . .

# Install dependencies
RUN npm install

# Expose the port the app runs on
EXPOSE 3040

# Command to run the start script
CMD ["sh", "start.sh"]
</file>

<file path="LICENSE">
GNU AFFERO GENERAL PUBLIC LICENSE
                       Version 3, 19 November 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.

                            Preamble

  The GNU Affero General Public License is a free, copyleft license for
software and other kinds of works, specifically designed to ensure
cooperation with the community in the case of network server software.

  The licenses for most software and other practical works are designed
to take away your freedom to share and change the works.  By contrast,
our General Public Licenses are intended to guarantee your freedom to
share and change all versions of a program--to make sure it remains free
software for all its users.

  When we speak of free software, we are referring to freedom, not
price.  Our General Public Licenses are designed to make sure that you
have the freedom to distribute copies of free software (and charge for
them if you wish), that you receive source code or can get it if you
want it, that you can change the software or use pieces of it in new
free programs, and that you know you can do these things.

  Developers that use our General Public Licenses protect your rights
with two steps: (1) assert copyright on the software, and (2) offer
you this License which gives you legal permission to copy, distribute
and/or modify the software.

  A secondary benefit of defending all users' freedom is that
improvements made in alternate versions of the program, if they
receive widespread use, become available for other developers to
incorporate.  Many developers of free software are heartened and
encouraged by the resulting cooperation.  However, in the case of
software used on network servers, this result may fail to come about.
The GNU General Public License permits making a modified version and
letting the public access it on a server without ever releasing its
source code to the public.

  The GNU Affero General Public License is designed specifically to
ensure that, in such cases, the modified source code becomes available
to the community.  It requires the operator of a network server to
provide the source code of the modified version running there to the
users of that server.  Therefore, public use of a modified version, on
a publicly accessible server, gives the public access to the source
code of the modified version.

  An older license, called the Affero General Public License and
published by Affero, was designed to accomplish similar goals.  This is
a different license, not a version of the Affero GPL, but Affero has
released a new version of the Affero GPL which permits relicensing under
this license.

  The precise terms and conditions for copying, distribution and
modification follow.

                       TERMS AND CONDITIONS

  0. Definitions.

  "This License" refers to version 3 of the GNU Affero General Public License.

  "Copyright" also means copyright-like laws that apply to other kinds of
works, such as semiconductor masks.

  "The Program" refers to any copyrightable work licensed under this
License.  Each licensee is addressed as "you".  "Licensees" and
"recipients" may be individuals or organizations.

  To "modify" a work means to copy from or adapt all or part of the work
in a fashion requiring copyright permission, other than the making of an
exact copy.  The resulting work is called a "modified version" of the
earlier work or a work "based on" the earlier work.

  A "covered work" means either the unmodified Program or a work based
on the Program.

  To "propagate" a work means to do anything with it that, without
permission, would make you directly or secondarily liable for
infringement under applicable copyright law, except executing it on a
computer or modifying a private copy.  Propagation includes copying,
distribution (with or without modification), making available to the
public, and in some countries other activities as well.

  To "convey" a work means any kind of propagation that enables other
parties to make or receive copies.  Mere interaction with a user through
a computer network, with no transfer of a copy, is not conveying.

  An interactive user interface displays "Appropriate Legal Notices"
to the extent that it includes a convenient and prominently visible
feature that (1) displays an appropriate copyright notice, and (2)
tells the user that there is no warranty for the work (except to the
extent that warranties are provided), that licensees may convey the
work under this License, and how to view a copy of this License.  If
the interface presents a list of user commands or options, such as a
menu, a prominent item in the list meets this criterion.

  1. Source Code.

  The "source code" for a work means the preferred form of the work
for making modifications to it.  "Object code" means any non-source
form of a work.

  A "Standard Interface" means an interface that either is an official
standard defined by a recognized standards body, or, in the case of
interfaces specified for a particular programming language, one that
is widely used among developers working in that language.

  The "System Libraries" of an executable work include anything, other
than the work as a whole, that (a) is included in the normal form of
packaging a Major Component, but which is not part of that Major
Component, and (b) serves only to enable use of the work with that
Major Component, or to implement a Standard Interface for which an
implementation is available to the public in source code form.  A
"Major Component", in this context, means a major essential component
(kernel, window system, and so on) of the specific operating system
(if any) on which the executable work runs, or a compiler used to
produce the work, or an object code interpreter used to run it.

  The "Corresponding Source" for a work in object code form means all
the source code needed to generate, install, and (for an executable
work) run the object code and to modify the work, including scripts to
control those activities.  However, it does not include the work's
System Libraries, or general-purpose tools or generally available free
programs which are used unmodified in performing those activities but
which are not part of the work.  For example, Corresponding Source
includes interface definition files associated with source files for
the work, and the source code for shared libraries and dynamically
linked subprograms that the work is specifically designed to require,
such as by intimate data communication or control flow between those
subprograms and other parts of the work.

  The Corresponding Source need not include anything that users
can regenerate automatically from other parts of the Corresponding
Source.

  The Corresponding Source for a work in source code form is that
same work.

  2. Basic Permissions.

  All rights granted under this License are granted for the term of
copyright on the Program, and are irrevocable provided the stated
conditions are met.  This License explicitly affirms your unlimited
permission to run the unmodified Program.  The output from running a
covered work is covered by this License only if the output, given its
content, constitutes a covered work.  This License acknowledges your
rights of fair use or other equivalent, as provided by copyright law.

  You may make, run and propagate covered works that you do not
convey, without conditions so long as your license otherwise remains
in force.  You may convey covered works to others for the sole purpose
of having them make modifications exclusively for you, or provide you
with facilities for running those works, provided that you comply with
the terms of this License in conveying all material for which you do
not control copyright.  Those thus making or running the covered works
for you must do so exclusively on your behalf, under your direction
and control, on terms that prohibit them from making any copies of
your copyrighted material outside their relationship with you.

  Conveying under any other circumstances is permitted solely under
the conditions stated below.  Sublicensing is not allowed; section 10
makes it unnecessary.

  3. Protecting Users' Legal Rights From Anti-Circumvention Law.

  No covered work shall be deemed part of an effective technological
measure under any applicable law fulfilling obligations under article
11 of the WIPO copyright treaty adopted on 20 December 1996, or
similar laws prohibiting or restricting circumvention of such
measures.

  When you convey a covered work, you waive any legal power to forbid
circumvention of technological measures to the extent such circumvention
is effected by exercising rights under this License with respect to
the covered work, and you disclaim any intention to limit operation or
modification of the work as a means of enforcing, against the work's
users, your or third parties' legal rights to forbid circumvention of
technological measures.

  4. Conveying Verbatim Copies.

  You may convey verbatim copies of the Program's source code as you
receive it, in any medium, provided that you conspicuously and
appropriately publish on each copy an appropriate copyright notice;
keep intact all notices stating that this License and any
non-permissive terms added in accord with section 7 apply to the code;
keep intact all notices of the absence of any warranty; and give all
recipients a copy of this License along with the Program.

  You may charge any price or no price for each copy that you convey,
and you may offer support or warranty protection for a fee.

  5. Conveying Modified Source Versions.

  You may convey a work based on the Program, or the modifications to
produce it from the Program, in the form of source code under the
terms of section 4, provided that you also meet all of these conditions:

    a) The work must carry prominent notices stating that you modified
    it, and giving a relevant date.

    b) The work must carry prominent notices stating that it is
    released under this License and any conditions added under section
    7.  This requirement modifies the requirement in section 4 to
    "keep intact all notices".

    c) You must license the entire work, as a whole, under this
    License to anyone who comes into possession of a copy.  This
    License will therefore apply, along with any applicable section 7
    additional terms, to the whole of the work, and all its parts,
    regardless of how they are packaged.  This License gives no
    permission to license the work in any other way, but it does not
    invalidate such permission if you have separately received it.

    d) If the work has interactive user interfaces, each must display
    Appropriate Legal Notices; however, if the Program has interactive
    interfaces that do not display Appropriate Legal Notices, your
    work need not make them do so.

  A compilation of a covered work with other separate and independent
works, which are not by their nature extensions of the covered work,
and which are not combined with it such as to form a larger program,
in or on a volume of a storage or distribution medium, is called an
"aggregate" if the compilation and its resulting copyright are not
used to limit the access or legal rights of the compilation's users
beyond what the individual works permit.  Inclusion of a covered work
in an aggregate does not cause this License to apply to the other
parts of the aggregate.

  6. Conveying Non-Source Forms.

  You may convey a covered work in object code form under the terms
of sections 4 and 5, provided that you also convey the
machine-readable Corresponding Source under the terms of this License,
in one of these ways:

    a) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by the
    Corresponding Source fixed on a durable physical medium
    customarily used for software interchange.

    b) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by a
    written offer, valid for at least three years and valid for as
    long as you offer spare parts or customer support for that product
    model, to give anyone who possesses the object code either (1) a
    copy of the Corresponding Source for all the software in the
    product that is covered by this License, on a durable physical
    medium customarily used for software interchange, for a price no
    more than your reasonable cost of physically performing this
    conveying of source, or (2) access to copy the
    Corresponding Source from a network server at no charge.

    c) Convey individual copies of the object code with a copy of the
    written offer to provide the Corresponding Source.  This
    alternative is allowed only occasionally and noncommercially, and
    only if you received the object code with such an offer, in accord
    with subsection 6b.

    d) Convey the object code by offering access from a designated
    place (gratis or for a charge), and offer equivalent access to the
    Corresponding Source in the same way through the same place at no
    further charge.  You need not require recipients to copy the
    Corresponding Source along with the object code.  If the place to
    copy the object code is a network server, the Corresponding Source
    may be on a different server (operated by you or a third party)
    that supports equivalent copying facilities, provided you maintain
    clear directions next to the object code saying where to find the
    Corresponding Source.  Regardless of what server hosts the
    Corresponding Source, you remain obligated to ensure that it is
    available for as long as needed to satisfy these requirements.

    e) Convey the object code using peer-to-peer transmission, provided
    you inform other peers where the object code and Corresponding
    Source of the work are being offered to the general public at no
    charge under subsection 6d.

  A separable portion of the object code, whose source code is excluded
from the Corresponding Source as a System Library, need not be
included in conveying the object code work.

  A "User Product" is either (1) a "consumer product", which means any
tangible personal property which is normally used for personal, family,
or household purposes, or (2) anything designed or sold for incorporation
into a dwelling.  In determining whether a product is a consumer product,
doubtful cases shall be resolved in favor of coverage.  For a particular
product received by a particular user, "normally used" refers to a
typical or common use of that class of product, regardless of the status
of the particular user or of the way in which the particular user
actually uses, or expects or is expected to use, the product.  A product
is a consumer product regardless of whether the product has substantial
commercial, industrial or non-consumer uses, unless such uses represent
the only significant mode of use of the product.

  "Installation Information" for a User Product means any methods,
procedures, authorization keys, or other information required to install
and execute modified versions of a covered work in that User Product from
a modified version of its Corresponding Source.  The information must
suffice to ensure that the continued functioning of the modified object
code is in no case prevented or interfered with solely because
modification has been made.

  If you convey an object code work under this section in, or with, or
specifically for use in, a User Product, and the conveying occurs as
part of a transaction in which the right of possession and use of the
User Product is transferred to the recipient in perpetuity or for a
fixed term (regardless of how the transaction is characterized), the
Corresponding Source conveyed under this section must be accompanied
by the Installation Information.  But this requirement does not apply
if neither you nor any third party retains the ability to install
modified object code on the User Product (for example, the work has
been installed in ROM).

  The requirement to provide Installation Information does not include a
requirement to continue to provide support service, warranty, or updates
for a work that has been modified or installed by the recipient, or for
the User Product in which it has been modified or installed.  Access to a
network may be denied when the modification itself materially and
adversely affects the operation of the network or violates the rules and
protocols for communication across the network.

  Corresponding Source conveyed, and Installation Information provided,
in accord with this section must be in a format that is publicly
documented (and with an implementation available to the public in
source code form), and must require no special password or key for
unpacking, reading or copying.

  7. Additional Terms.

  "Additional permissions" are terms that supplement the terms of this
License by making exceptions from one or more of its conditions.
Additional permissions that are applicable to the entire Program shall
be treated as though they were included in this License, to the extent
that they are valid under applicable law.  If additional permissions
apply only to part of the Program, that part may be used separately
under those permissions, but the entire Program remains governed by
this License without regard to the additional permissions.

  When you convey a copy of a covered work, you may at your option
remove any additional permissions from that copy, or from any part of
it.  (Additional permissions may be written to require their own
removal in certain cases when you modify the work.)  You may place
additional permissions on material, added by you to a covered work,
for which you have or can give appropriate copyright permission.

  Notwithstanding any other provision of this License, for material you
add to a covered work, you may (if authorized by the copyright holders of
that material) supplement the terms of this License with terms:

    a) Disclaiming warranty or limiting liability differently from the
    terms of sections 15 and 16 of this License; or

    b) Requiring preservation of specified reasonable legal notices or
    author attributions in that material or in the Appropriate Legal
    Notices displayed by works containing it; or

    c) Prohibiting misrepresentation of the origin of that material, or
    requiring that modified versions of such material be marked in
    reasonable ways as different from the original version; or

    d) Limiting the use for publicity purposes of names of licensors or
    authors of the material; or

    e) Declining to grant rights under trademark law for use of some
    trade names, trademarks, or service marks; or

    f) Requiring indemnification of licensors and authors of that
    material by anyone who conveys the material (or modified versions of
    it) with contractual assumptions of liability to the recipient, for
    any liability that these contractual assumptions directly impose on
    those licensors and authors.

  All other non-permissive additional terms are considered "further
restrictions" within the meaning of section 10.  If the Program as you
received it, or any part of it, contains a notice stating that it is
governed by this License along with a term that is a further
restriction, you may remove that term.  If a license document contains
a further restriction but permits relicensing or conveying under this
License, you may add to a covered work material governed by the terms
of that license document, provided that the further restriction does
not survive such relicensing or conveying.

  If you add terms to a covered work in accord with this section, you
must place, in the relevant source files, a statement of the
additional terms that apply to those files, or a notice indicating
where to find the applicable terms.

  Additional terms, permissive or non-permissive, may be stated in the
form of a separately written license, or stated as exceptions;
the above requirements apply either way.

  8. Termination.

  You may not propagate or modify a covered work except as expressly
provided under this License.  Any attempt otherwise to propagate or
modify it is void, and will automatically terminate your rights under
this License (including any patent licenses granted under the third
paragraph of section 11).

  However, if you cease all violation of this License, then your
license from a particular copyright holder is reinstated (a)
provisionally, unless and until the copyright holder explicitly and
finally terminates your license, and (b) permanently, if the copyright
holder fails to notify you of the violation by some reasonable means
prior to 60 days after the cessation.

  Moreover, your license from a particular copyright holder is
reinstated permanently if the copyright holder notifies you of the
violation by some reasonable means, this is the first time you have
received notice of violation of this License (for any work) from that
copyright holder, and you cure the violation prior to 30 days after
your receipt of the notice.

  Termination of your rights under this section does not terminate the
licenses of parties who have received copies or rights from you under
this License.  If your rights have been terminated and not permanently
reinstated, you do not qualify to receive new licenses for the same
material under section 10.

  9. Acceptance Not Required for Having Copies.

  You are not required to accept this License in order to receive or
run a copy of the Program.  Ancillary propagation of a covered work
occurring solely as a consequence of using peer-to-peer transmission
to receive a copy likewise does not require acceptance.  However,
nothing other than this License grants you permission to propagate or
modify any covered work.  These actions infringe copyright if you do
not accept this License.  Therefore, by modifying or propagating a
covered work, you indicate your acceptance of this License to do so.

  10. Automatic Licensing of Downstream Recipients.

  Each time you convey a covered work, the recipient automatically
receives a license from the original licensors, to run, modify and
propagate that work, subject to this License.  You are not responsible
for enforcing compliance by third parties with this License.

  An "entity transaction" is a transaction transferring control of an
organization, or substantially all assets of one, or subdividing an
organization, or merging organizations.  If propagation of a covered
work results from an entity transaction, each party to that
transaction who receives a copy of the work also receives whatever
licenses to the work the party's predecessor in interest had or could
give under the previous paragraph, plus a right to possession of the
Corresponding Source of the work from the predecessor in interest, if
the predecessor has it or can get it with reasonable efforts.

  You may not impose any further restrictions on the exercise of the
rights granted or affirmed under this License.  For example, you may
not impose a license fee, royalty, or other charge for exercise of
rights granted under this License, and you may not initiate litigation
(including a cross-claim or counterclaim in a lawsuit) alleging that
any patent claim is infringed by making, using, selling, offering for
sale, or importing the Program or any portion of it.

  11. Patents.

  A "contributor" is a copyright holder who authorizes use under this
License of the Program or a work on which the Program is based.  The
work thus licensed is called the contributor's "contributor version".

  A contributor's "essential patent claims" are all patent claims
owned or controlled by the contributor, whether already acquired or
hereafter acquired, that would be infringed by some manner, permitted
by this License, of making, using, or selling its contributor version,
but do not include claims that would be infringed only as a
consequence of further modification of the contributor version.  For
purposes of this definition, "control" includes the right to grant
patent sublicenses in a manner consistent with the requirements of
this License.

  Each contributor grants you a non-exclusive, worldwide, royalty-free
patent license under the contributor's essential patent claims, to
make, use, sell, offer for sale, import and otherwise run, modify and
propagate the contents of its contributor version.

  In the following three paragraphs, a "patent license" is any express
agreement or commitment, however denominated, not to enforce a patent
(such as an express permission to practice a patent or covenant not to
sue for patent infringement).  To "grant" such a patent license to a
party means to make such an agreement or commitment not to enforce a
patent against the party.

  If you convey a covered work, knowingly relying on a patent license,
and the Corresponding Source of the work is not available for anyone
to copy, free of charge and under the terms of this License, through a
publicly available network server or other readily accessible means,
then you must either (1) cause the Corresponding Source to be so
available, or (2) arrange to deprive yourself of the benefit of the
patent license for this particular work, or (3) arrange, in a manner
consistent with the requirements of this License, to extend the patent
license to downstream recipients.  "Knowingly relying" means you have
actual knowledge that, but for the patent license, your conveying the
covered work in a country, or your recipient's use of the covered work
in a country, would infringe one or more identifiable patents in that
country that you have reason to believe are valid.

  If, pursuant to or in connection with a single transaction or
arrangement, you convey, or propagate by procuring conveyance of, a
covered work, and grant a patent license to some of the parties
receiving the covered work authorizing them to use, propagate, modify
or convey a specific copy of the covered work, then the patent license
you grant is automatically extended to all recipients of the covered
work and works based on it.

  A patent license is "discriminatory" if it does not include within
the scope of its coverage, prohibits the exercise of, or is
conditioned on the non-exercise of one or more of the rights that are
specifically granted under this License.  You may not convey a covered
work if you are a party to an arrangement with a third party that is
in the business of distributing software, under which you make payment
to the third party based on the extent of your activity of conveying
the work, and under which the third party grants, to any of the
parties who would receive the covered work from you, a discriminatory
patent license (a) in connection with copies of the covered work
conveyed by you (or copies made from those copies), or (b) primarily
for and in connection with specific products or compilations that
contain the covered work, unless you entered into that arrangement,
or that patent license was granted, prior to 28 March 2007.

  Nothing in this License shall be construed as excluding or limiting
any implied license or other defenses to infringement that may
otherwise be available to you under applicable patent law.

  12. No Surrender of Others' Freedom.

  If conditions are imposed on you (whether by court order, agreement or
otherwise) that contradict the conditions of this License, they do not
excuse you from the conditions of this License.  If you cannot convey a
covered work so as to satisfy simultaneously your obligations under this
License and any other pertinent obligations, then as a consequence you may
not convey it at all.  For example, if you agree to terms that obligate you
to collect a royalty for further conveying from those to whom you convey
the Program, the only way you could satisfy both those terms and this
License would be to refrain entirely from conveying the Program.

  13. Remote Network Interaction; Use with the GNU General Public License.

  Notwithstanding any other provision of this License, if you modify the
Program, your modified version must prominently offer all users
interacting with it remotely through a computer network (if your version
supports such interaction) an opportunity to receive the Corresponding
Source of your version by providing access to the Corresponding Source
from a network server at no charge, through some standard or customary
means of facilitating copying of software.  This Corresponding Source
shall include the Corresponding Source for any work covered by version 3
of the GNU General Public License that is incorporated pursuant to the
following paragraph.

  Notwithstanding any other provision of this License, you have
permission to link or combine any covered work with a work licensed
under version 3 of the GNU General Public License into a single
combined work, and to convey the resulting work.  The terms of this
License will continue to apply to the part which is the covered work,
but the work with which it is combined will remain governed by version
3 of the GNU General Public License.

  14. Revised Versions of this License.

  The Free Software Foundation may publish revised and/or new versions of
the GNU Affero General Public License from time to time.  Such new versions
will be similar in spirit to the present version, but may differ in detail to
address new problems or concerns.

  Each version is given a distinguishing version number.  If the
Program specifies that a certain numbered version of the GNU Affero General
Public License "or any later version" applies to it, you have the
option of following the terms and conditions either of that numbered
version or of any later version published by the Free Software
Foundation.  If the Program does not specify a version number of the
GNU Affero General Public License, you may choose any version ever published
by the Free Software Foundation.

  If the Program specifies that a proxy can decide which future
versions of the GNU Affero General Public License can be used, that proxy's
public statement of acceptance of a version permanently authorizes you
to choose that version for the Program.

  Later license versions may give you additional or different
permissions.  However, no additional obligations are imposed on any
author or copyright holder as a result of your choosing to follow a
later version.

  15. Disclaimer of Warranty.

  THERE IS NO WARRANTY FOR THE PROGRAM, TO THE EXTENT PERMITTED BY
APPLICABLE LAW.  EXCEPT WHEN OTHERWISE STATED IN WRITING THE COPYRIGHT
HOLDERS AND/OR OTHER PARTIES PROVIDE THE PROGRAM "AS IS" WITHOUT WARRANTY
OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING, BUT NOT LIMITED TO,
THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
PURPOSE.  THE ENTIRE RISK AS TO THE QUALITY AND PERFORMANCE OF THE PROGRAM
IS WITH YOU.  SHOULD THE PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF
ALL NECESSARY SERVICING, REPAIR OR CORRECTION.

  16. Limitation of Liability.

  IN NO EVENT UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING
WILL ANY COPYRIGHT HOLDER, OR ANY OTHER PARTY WHO MODIFIES AND/OR CONVEYS
THE PROGRAM AS PERMITTED ABOVE, BE LIABLE TO YOU FOR DAMAGES, INCLUDING ANY
GENERAL, SPECIAL, INCIDENTAL OR CONSEQUENTIAL DAMAGES ARISING OUT OF THE
USE OR INABILITY TO USE THE PROGRAM (INCLUDING BUT NOT LIMITED TO LOSS OF
DATA OR DATA BEING RENDERED INACCURATE OR LOSSES SUSTAINED BY YOU OR THIRD
PARTIES OR A FAILURE OF THE PROGRAM TO OPERATE WITH ANY OTHER PROGRAMS),
EVEN IF SUCH HOLDER OR OTHER PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF
SUCH DAMAGES.

  17. Interpretation of Sections 15 and 16.

  If the disclaimer of warranty and limitation of liability provided
above cannot be given local legal effect according to their terms,
reviewing courts shall apply local law that most closely approximates
an absolute waiver of all civil liability in connection with the
Program, unless a warranty or assumption of liability accompanies a
copy of the Program in return for a fee.

                     END OF TERMS AND CONDITIONS

            How to Apply These Terms to Your New Programs

  If you develop a new program, and you want it to be of the greatest
possible use to the public, the best way to achieve this is to make it
free software which everyone can redistribute and change under these terms.

  To do so, attach the following notices to the program.  It is safest
to attach them to the start of each source file to most effectively
state the exclusion of warranty; and each file should have at least
the "copyright" line and a pointer to where the full notice is found.

    <one line to give the program's name and a brief idea of what it does.>
    Copyright (C) <year>  <name of author>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

Also add information on how to contact you by electronic and paper mail.

  If your software can interact with users remotely through a computer
network, you should also make sure that it provides a way for users to
get its source.  For example, if your program is a web application, its
interface could display a "Source" link that leads users to an archive
of the code.  There are many ways you could offer source, and different
solutions will be better for different programs; see section 13 for the
specific requirements.

  You should also get your employer (if you work as a programmer) or school,
if any, to sign a "copyright disclaimer" for the program, if necessary.
For more information on this, and how to apply and follow the GNU AGPL, see
<https://www.gnu.org/licenses/>.
</file>

<file path="package.json">
{
  "name": "chatgpt",
  "version": "1.0.0",
  "description": "OpenAI API Free Reverse Proxy",
  "type": "module",
  "main": "app.js",
  "scripts": {
    "start": "tsc && node dist/app.js",
    "watch": "tsc-watch --onSuccess \"node dist/app.js\"",
    "build": "tsc"
  },
  "author": "Pawan Osman",
  "license": "AGPL-3.0",
  "repository": {
    "type": "git",
    "url": "https://github.com/PawanOsman/ChatGPT.git"
  },
  "dependencies": {
    "axios": "^1.6.7",
    "body-parser": "^1.20.2",
    "dotenv": "^16.4.5",
    "express": "^4.18.3",
    "gpt-3-encoder": "^1.1.4"
  },
  "devDependencies": {
    "@types/express": "^4.17.21",
    "ts-node": "^10.9.2",
    "typescript": "^5.3.3"
  }
}
</file>

<file path="src/app.ts">
import express, { Request, Response, NextFunction } from "express";
import { ChildProcessWithoutNullStreams, spawn } from "child_process";
import fs from "fs";
import path from "path";
import bodyParser from "body-parser";
import axios from "axios";
import https from "https";
import os from "os";
import { encode } from "gpt-3-encoder";
import { randomUUID, randomInt, createHash } from "crypto";
import { config } from "dotenv";

config();

// Constants for the server and API configuration
const port = process.env.SERVER_PORT || 3040;
const baseUrl = "https://chat.openai.com";
const apiUrl = `${baseUrl}/backend-anon/conversation`;
const refreshInterval = 60000; // Interval to refresh token in ms
const errorWait = 120000; // Wait time in ms after an error
const newSessionRetries: number =
  parseInt(process.env.NEW_SESSION_RETRIES) || 5; // Number of retries to get a new session
const userAgent =
  process.env.USER_AGENT ||
  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36";
const authKey: string =
  process.env.API_KEY || null;  // Authorized client apiKey

let cloudflared: ChildProcessWithoutNullStreams;

// Type definition for the session object
type Session = {
  deviceId: string;
  persona: string;
  arkose: {
    required: boolean;
    dx: any;
  };
  turnstile: {
    required: boolean;
  };
  proofofwork: {
    required: boolean;
    seed: string;
    difficulty: string;
  };
  token: string;
};

// Function to wait for a specified duration
const wait = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

function GenerateCompletionId(prefix: string = "cmpl-") {
  const characters =
    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  const length = 28;

  for (let i = 0; i < length; i++) {
    prefix += characters.charAt(Math.floor(Math.random() * characters.length));
  }

  return prefix;
}

async function* chunksToLines(chunksAsync: any) {
  let previous = "";
  for await (const chunk of chunksAsync) {
    const bufferChunk = Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk);
    previous += bufferChunk;
    let eolIndex: number;
    while ((eolIndex = previous.indexOf("\n")) >= 0) {
      // line includes the EOL
      const line = previous.slice(0, eolIndex + 1).trimEnd();
      if (line === "data: [DONE]") break;
      if (line.startsWith("data: ")) yield line;
      previous = previous.slice(eolIndex + 1);
    }
  }
}

async function* linesToMessages(linesAsync: any) {
  for await (const line of linesAsync) {
    const message = line.substring("data :".length);

    yield message;
  }
}

async function* StreamCompletion(data: any) {
  yield* linesToMessages(chunksToLines(data));
}

// Setup axios instance for API requests with predefined configurations
const axiosInstance = axios.create({
  httpsAgent: new https.Agent({ rejectUnauthorized: false }),
  proxy:
    process.env.PROXY === "true"
      ? {
          host: process.env.PROXY_HOST,
          port: Number(process.env.PROXY_PORT),
          auth:
            process.env.PROXY_AUTH === "true"
              ? {
                  username: process.env.PROXY_USERNAME,
                  password: process.env.PROXY_PASSWORD,
                }
              : undefined,
          protocol: process.env.PROXY_PROTOCOL,
        }
      : false,
  headers: {
    accept: "*/*",
    "accept-language": "en-US,en;q=0.9",
    "cache-control": "no-cache",
    "content-type": "application/json",
    "oai-language": "en-US",
    origin: baseUrl,
    pragma: "no-cache",
    referer: baseUrl,
    "sec-ch-ua":
      '"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"Windows"',
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "user-agent": userAgent,
  },
});

// Generate a proof token for the OpenAI API
function GenerateProofToken(
  seed: string,
  diff: string,
  userAgent: string
): string {
  const cores: number[] = [8, 12, 16, 24];
  const screens: number[] = [3000, 4000, 6000];

  const core = cores[randomInt(0, cores.length)];
  const screen = screens[randomInt(0, screens.length)];

  const now = new Date(Date.now() - 8 * 3600 * 1000);
  const parseTime = now.toUTCString().replace("GMT", "GMT-0500 (Eastern Time)");

  const config = [core + screen, parseTime, 4294705152, 0, userAgent];

  const diffLen = diff.length / 2;

  for (let i = 0; i < 100000; i++) {
    config[3] = i;
    const jsonData = JSON.stringify(config);
    const base = Buffer.from(jsonData).toString("base64");
    const hashValue = createHash("sha3-512")
      .update(seed + base)
      .digest();

    if (hashValue.toString("hex").substring(0, diffLen) <= diff) {
      const result = "gAAAAAB" + base;
      return result;
    }
  }

  const fallbackBase = Buffer.from(`"${seed}"`).toString("base64");
  return "gAAAAABwQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D" + fallbackBase;
}

// Function to get a new session ID and token from the OpenAI API
async function getNewSession(retries: number = 0): Promise<Session> {
  let newDeviceId = randomUUID();
  try {
    const response = await axiosInstance.post(
      `${baseUrl}/backend-anon/sentinel/chat-requirements`,
      {},
      {
        headers: { "oai-device-id": newDeviceId },
      }
    );

    let session: Session = response.data as Session;
    session.deviceId = newDeviceId;

    return session;
  } catch (error) {
    await wait(500);
    return retries < newSessionRetries ? getNewSession(retries + 1) : null;
  }
}

// Middleware to enable CORS and handle pre-flight requests
function enableCORS(req: Request, res: Response, next: NextFunction) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "*");
  res.header("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
  if (req.method === "OPTIONS") {
    return res.status(200).end();
  }
  next();
}

// Middleware to handle chat completions
async function handleChatCompletion(req: Request, res: Response) {
  // If .env sets API_KEY and is not empty, the apiKey of req.headers will be verified.
  if (authKey) {
    const clientApiKey = req.headers.authorization?.split(' ')[1] ?? "null";
    if (!clientApiKey || clientApiKey != authKey) {
      console.log(
        "Request:",
        `${req.method} ${req.originalUrl}`,
        `${req.body?.messages?.length ?? 0} messages`,
        `ClientKey: ${clientApiKey} Verify Failed!`
      );

      res.write(
        JSON.stringify({
          status: false,
          error: {
            message: `Incorrect API key provided: ${clientApiKey}, Authorized access only!`,
            type: "invalid_request_error",
            code: "invalid_api_key"
          },
          support: "https://discord.pawan.krd",
        })
      );
      return res.end();
    }
  }

  console.log(
    "Request:",
    `${req.method} ${req.originalUrl}`,
    `${req.body?.messages?.length ?? 0} messages`,
    req.body.stream ? "(stream-enabled)" : "(stream-disabled)"
  );
  try {
    let session = await getNewSession();

    if (!session) {
      console.error("Error getting a new session...");
      console.error("If this error persists, your country may not be supported yet.");
      console.error("If your country was the issue, please consider using a U.S. VPN or a U.S. residential proxy.");
      res.write(
        JSON.stringify({
          status: false,
          error: {
            message: `Error getting a new session, If this error persists, your country may not be supported yet. If your country was the issue, please consider using a U.S. VPN or a U.S. residential proxy.`,
            type: "invalid_request_error",
          },
          support: "https://discord.pawan.krd",
        })
      );

      return res.end();
    }

    let proofToken = GenerateProofToken(
      session.proofofwork.seed,
      session.proofofwork.difficulty,
      userAgent
    );

    const body = {
      action: "next",
      messages: req.body.messages.map(
        (message: { role: any; content: any }) => ({
          author: { role: message.role },
          content: { content_type: "text", parts: [message.content] },
        })
      ),
      parent_message_id: randomUUID(),
      model: "text-davinci-002-render-sha",
      timezone_offset_min: -180,
      suggestions: [],
      history_and_training_disabled: true,
      conversation_mode: { kind: "primary_assistant" },
      websocket_request_id: randomUUID(),
    };

    let promptTokens = 0;
    let completionTokens = 0;

    for (let message of req.body.messages) {
      promptTokens += encode(message.content).length;
    }

    const response = await axiosInstance.post(apiUrl, body, {
      responseType: "stream",
      headers: {
        "oai-device-id": session.deviceId,
        "openai-sentinel-chat-requirements-token": session.token,
        "openai-sentinel-proof-token": proofToken,
      },
    });

    // Set the response headers based on the request type
    if (req.body.stream) {
      res.setHeader("Content-Type", "text/event-stream");
      res.setHeader("Cache-Control", "no-cache");
      res.setHeader("Connection", "keep-alive");
    } else {
      res.setHeader("Content-Type", "application/json");
    }

    let fullContent = "";
    let requestId = GenerateCompletionId("chatcmpl-");
    let created = Math.floor(Date.now() / 1000); // Unix timestamp in seconds
    let finish_reason = null;
    let error: string;

    for await (const message of StreamCompletion(response.data)) {
      // Skip heartbeat detection
      if (message.match(/^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{6}$/))
        continue;

      const parsed = JSON.parse(message);

      if (parsed.error) {
        error = `Error message from OpenAI: ${parsed.error}`;
        finish_reason = "stop";
        break;
      }

      let content = parsed?.message?.content?.parts[0] ?? "";
      let status = parsed?.message?.status ?? "";

      for (let message of req.body.messages) {
        if (message.content === content) {
          content = "";
          break;
        }
      }

      switch (status) {
        case "in_progress":
          finish_reason = null;
          break;
        case "finished_successfully":
          let finish_reason_data =
            parsed?.message?.metadata?.finish_details?.type ?? null;
          switch (finish_reason_data) {
            case "max_tokens":
              finish_reason = "length";
              break;
            case "stop":
            default:
              finish_reason = "stop";
          }
          break;
        default:
          finish_reason = null;
      }

      if (content === "") continue;

      let completionChunk = content.replace(fullContent, "");

      completionTokens += encode(completionChunk).length;

      if (req.body.stream) {
        let response = {
          id: requestId,
          created: created,
          object: "chat.completion.chunk",
          model: "gpt-3.5-turbo",
          choices: [
            {
              delta: {
                content: completionChunk,
              },
              index: 0,
              finish_reason: finish_reason,
            },
          ],
        };

        res.write(`data: ${JSON.stringify(response)}\n\n`);
      }

      fullContent = content.length > fullContent.length ? content : fullContent;
    }

    if (req.body.stream) {
      res.write(
        `data: ${JSON.stringify({
          id: requestId,
          created: created,
          object: "chat.completion.chunk",
          model: "gpt-3.5-turbo",
          choices: [
            {
              delta: {
                content: error ?? "",
              },
              index: 0,
              finish_reason: finish_reason,
            },
          ],
        })}\n\n`
      );
    } else {
      res.write(
        JSON.stringify({
          id: requestId,
          created: created,
          model: "gpt-3.5-turbo",
          object: "chat.completion",
          choices: [
            {
              finish_reason: finish_reason,
              index: 0,
              message: {
                content: error ?? fullContent,
                role: "assistant",
              },
            },
          ],
          usage: {
            prompt_tokens: promptTokens,
            completion_tokens: completionTokens,
            total_tokens: promptTokens + completionTokens,
          },
        })
      );
    }

    res.end();
  } catch (error: any) {
    // console.log("Error:", error.response?.data ?? error.message);
    if (!res.headersSent) res.setHeader("Content-Type", "application/json");
    // console.error("Error handling chat completion:", error);
    res.write(
      JSON.stringify({
        status: false,
        error: {
          message:
            "An error occurred. please try again. Additionally, ensure that your request complies with OpenAI's policy.",
          type: "invalid_request_error",
        },
        support: "https://discord.pawan.krd",
      })
    );
    res.end();
  }
}

// Initialize Express app and use middlewares
const app = express();
app.use(bodyParser.json());
app.use(enableCORS);

// Route to handle POST requests for chat completions
app.post("/v1/chat/completions", handleChatCompletion);

// 404 handler for unmatched routes
app.use((req, res) =>
  res.status(404).send({
    status: false,
    error: {
      message: `The requested endpoint (${req.method.toLocaleUpperCase()} ${
        req.path
      }) was not found. please make sure to use "http://localhost:3040/v1" as the base URL.`,
      type: "invalid_request_error",
    },
    support: "https://discord.pawan.krd",
  })
);

async function DownloadCloudflared(): Promise<string> {
  const platform = os.platform();
  let url: string;

  if (platform === "win32") {
    const arch = os.arch() === "x64" ? "amd64" : "386";
    url = `https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-windows-${arch}.exe`;
  } else {
    let arch = os.arch();
    switch (arch) {
      case "x64":
        arch = "amd64";
        break;
      case "arm":
      case "arm64":
        break;
      default:
        arch = "amd64"; // Default to amd64 if unknown architecture
    }
    const platformLower = platform.toLowerCase();
    url = `https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-${platformLower}-${arch}`;
  }

  const fileName = platform === "win32" ? "cloudflared.exe" : "cloudflared";
  const filePath = path.resolve(fileName);

  if (fs.existsSync(filePath)) {
    return filePath;
  }

  try {
    const response = await axiosInstance({
      method: "get",
      url: url,
      responseType: "stream",
    });

    const writer = fs.createWriteStream(filePath);

    response.data.pipe(writer);

    return new Promise<string>((resolve, reject) => {
      writer.on("finish", () => {
        if (platform !== "win32") {
          fs.chmodSync(filePath, 0o755);
        }
        resolve(filePath);
      });

      writer.on("error", reject);
    });
  } catch (error: any) {
    // console.error("Failed to download file:", error.message);
    return null;
  }
}

async function StartCloudflaredTunnel(
  cloudflaredPath: string
): Promise<string> {
  if (!cloudflaredPath) {
    console.error("Failed to download Cloudflared executable.");
    return null;
  }

  const localUrl = `http://localhost:${port}`;
  return new Promise<string>((resolve, reject) => {
    cloudflared = spawn(cloudflaredPath, ["tunnel", "--url", localUrl]);

    cloudflared.stdout.on("data", (data: any) => {
      const output = data.toString();
      // console.log("Cloudflared Output:", output);

      // Adjusted regex to specifically match URLs that end with .trycloudflare.com
      const urlMatch = output.match(/https:\/\/[^\s]+\.trycloudflare\.com/);
      if (urlMatch) {
        let url = urlMatch[0];
        resolve(url);
      }
    });

    cloudflared.stderr.on("data", (data: any) => {
      const output = data.toString();
      // console.error("Error from cloudflared:", output);

      const urlMatch = output.match(/https:\/\/[^\s]+\.trycloudflare\.com/);
      if (urlMatch) {
        let url = urlMatch[0];
        resolve(url);
      }
    });

    cloudflared.on("close", (code: any) => {
      resolve(null);
      // console.log(`Cloudflared tunnel process exited with code ${code}`);
    });
  });
}

// Start the server and the session ID refresh loop
app.listen(port, async () => {
  if (process.env.CLOUDFLARED === undefined) process.env.CLOUDFLARED = "true";
  let cloudflared = process.env.CLOUDFLARED === "true";
  let filePath: string;
  let publicURL: string;
  if (cloudflared) {
    filePath = await DownloadCloudflared();
    publicURL = await StartCloudflaredTunnel(filePath);
  }

  console.log(`💡 Server is running at http://localhost:${port}`);
  console.log();
  console.log(`🔗 Local Base URL: http://localhost:${port}/v1`);
  console.log(
    `🔗 Local Endpoint: http://localhost:${port}/v1/chat/completions`
  );
  console.log();
  if (cloudflared && publicURL)
    console.log(`🔗 Public Base URL: ${publicURL}/v1`);
  if (cloudflared && publicURL)
    console.log(`🔗 Public Endpoint: ${publicURL}/v1/chat/completions`);
  else if (cloudflared && !publicURL) {
    console.log(
      "🔗 Public Endpoint: (Failed to start cloudflared tunnel, please restart the server.)"
    );
    if (filePath) fs.unlinkSync(filePath);
  }
  if (cloudflared && publicURL) console.log();
  console.log("📝 Author: Pawan.Krd");
  console.log(`🌐 Discord server: https://discord.gg/pawan`);
  console.log("🌍 GitHub Repository: https://github.com/PawanOsman/ChatGPT");
  console.log(
    `💖 Don't forget to star the repository if you like this project!`
  );
});
</file>

<file path="start.bat">
@echo off

IF NOT EXIST node_modules (
    echo Installing npm packages...
    call npm install
)

cls
echo Starting the application...
call npm start
</file>

<file path="start.sh">
#!/bin/sh

if [ ! -d "node_modules" ]; then
    echo "Installing npm packages..."
    npm install
fi

clear
echo "Starting the application..."
npm start
</file>

<file path="tsconfig.json">
{
	"compilerOptions": {
		"target": "es2022",
		"module": "ESNext",
		"strict": true,
		"declaration": true,
		"removeComments": true,
		"emitDecoratorMetadata": true,
		"experimentalDecorators": true,
		"allowSyntheticDefaultImports": true,
		"sourceMap": true,
		"outDir": "./dist",
		"baseUrl": "./",
		"paths": {
			"@/*": ["src/*"]
		},
		"incremental": true,
		"moduleResolution": "node",
		"skipLibCheck": true,
		"strictNullChecks": false,
		"noImplicitAny": false,
		"strictBindCallApply": false,
		"forceConsistentCasingInFileNames": false,
		"noFallthroughCasesInSwitch": false
	},
	"include": ["src"],
	"exclude": ["node_modules", "dist"]
}
</file>

<file path=".gitignore">
.vscode/
.idea/
.DS_Store
npm-debug.log
yarn.lock
yarn-error.log
node_modules/
dist/
*.tsbuildinfo
*.js
*.js.map
.parcel-cache
db.json
.env
cloudflared
</file>

<file path="README.md">
# This project is a bit outdated and isn’t working right now. We’ll update it, but we’re tied up with another project at the moment. In the meantime, you’re welcome to use [our hosted models for free](#accessing-our-hosted-api).


## ChatGPT `gpt-3.5-turbo` API for Free (as a Reverse Proxy)

Welcome to the ChatGPT API Free Reverse Proxy, offering free self-hosted API access to ChatGPT (`gpt-3.5-turbo`) with OpenAI's familiar structure, so no code changes are needed.

## Quick Links

- [Join our Discord Community](https://discord.pawan.krd) for support and questions.
	- ⚡Note: Your Discord account must be at least 7 days old to be able join our Discord community.

## Table of Contents

- [Features](#features)
- Option 1: [Installing/Self-Hosting Guide](#installingself-hosting-guide) (Without using any API key)
  - Method 1: [Using Docker](#using-docker) or [Run it with a Chat Web UI using docker-compose](./docker-compose)
  - Method 2: [Your PC/Server](#your-pcserver) (manually)
  - Method 3: [Termux on Android Phones](#termux-on-android-phones)
- Option 2: [Accessing Our Hosted API](#accessing-our-hosted-api) (Free)
- [Usage Examples](#usage-examples)
- [License](#license)

## Features

- **Streaming Response**: The API supports streaming response, so you can get the response as soon as it's available.
- **API Endpoint Compatibility**: Full alignment with official OpenAI API endpoints, ensuring hassle-free integration with existing OpenAI libraries.
- **Complimentary Access**: No charges for API usage, making advanced AI accessible to everyone even **without an API key**.

## Installing/Self-Hosting Guide

### Using Docker
1. Ensure Docker is installed by referring to the [Docker Installation Docs](https://docs.docker.com/engine/install/).
2. Run the following command:
   ```bash
   docker run -dp 3040:3040 pawanosman/chatgpt:latest
   ```
3. Done! You can now connect to your local server's API at:
   ```
   http://localhost:3040/v1/chat/completions
   ```
   Note that the base URL is `http://localhost:3040/v1`.

### Install with chat web interfaces
✅ You can run third-party chat web interfaces, such as BetterChatGPT and LobeChat, with this API using Docker Compose. [Click here for the installation guide](./docker-compose).

### Your PC/Server

To install and run the ChatGPT API Reverse Proxy on your PC/Server by following these steps:

Note: This option is not available to all countries yet. if you are from a country that is not supported, you can use a **U.S. VPN** or use **our hosted API**.

1. Ensure NodeJs (v19+) is installed: [Download NodeJs](https://nodejs.org/en/download)
2. Clone this repository:
   ```bash
   git clone https://github.com/PawanOsman/ChatGPT.git
   ```
3. Open `start.bat` (Windows) or `start.sh` (Linux with `bash start.sh` command) to install dependencies and launch the server.
4. Done, you can connect to your local server's API at:
   ```
   http://localhost:3040/v1/chat/completions
   ```
   Note that the base url will be `http://localhost:3040/v1`

To include installation instructions for Termux on Android devices, you can add the following section right after the instructions for Linux in the **Installing/Self-Hosting Guide**:

### Termux on Android Phones

To install and run the ChatGPT API Reverse Proxy on Android using Termux, follow these steps:

1. Install [Termux](https://play.google.com/store/apps/details?id=com.termux) from the Play Store.
2. Update Termux packages:
   ```bash
   apt update
   ```
3. Upgrade Termux packages:
   ```bash
   apt upgrade
   ```
4. Install git, Node.js, and npm:
   ```bash
   apt install -y git nodejs
   ```
5. Clone the repository:
   ```bash
   git clone https://github.com/PawanOsman/ChatGPT.git
   ```
6. Navigate to the cloned directory:
   ```bash
   cd ChatGPT
   ```
7. Start the server with:

   ```bash
   bash start.sh
   ```

8. Your local server will now be running and accessible at:

   ```
   http://localhost:3040/v1/chat/completions
   ```

   Note that the base url will be `http://localhost:3040/v1`

   You can now use this address to connect to your self-hosted ChatGPT API Reverse Proxy from Android applications/websites that support reverse proxy configurations, on the same device.

## Accessing Our Hosted API

Utilize our pre-hosted ChatGPT-like API for free by:

1. Joining our [Discord server](https://discord.pawan.krd).
2. Obtaining an API key from the `#Bot` channel with the `/key` command.
3. Incorporating the API key into your requests to:
   ```
   https://api.pawan.krd/v1/chat/completions
   ```

## Usage Examples

Leverage the same integration code as OpenAI's official libraries by simply adjusting the API key and base URL in your requests. For self-hosted setups, ensure to switch the base URL to your local server's address as mentioned above.

### Example Usage with OpenAI Libraries

#### Python Example

```python
import openai

openai.api_key = 'anything'
openai.base_url = "http://localhost:3040/v1/"

completion = openai.chat.completions.create(
    model="gpt-3.5-turbo",
    messages=[
        {"role": "user", "content": "How do I list all files in a directory using Python?"},
    ],
)

print(completion.choices[0].message.content)
```

#### Node.js Example

```js
import OpenAI from 'openai';

const openai = new OpenAI({
	apiKey: "anything",
	baseURL: "http://localhost:3040/v1",
});

const chatCompletion = await openai.chat.completions.create({
  messages: [{ role: 'user', content: 'Say this is a test' }],
  model: 'gpt-3.5-turbo',
});

console.log(chatCompletion.choices[0].message.content);
```

## License

This project is under the AGPL-3.0 License. Refer to the [LICENSE](LICENSE) file for detailed information.
</file>

</files>
