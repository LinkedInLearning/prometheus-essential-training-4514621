# Prometheus Essential Training
This is the repository for the LinkedIn Learning course `Prometheus Essential Training`. The full course is available from [LinkedIn Learning][lil-course-url].

![lil-thumbnail-url]

<p>Prometheus is one of the most widely used, open-source monitoring systems in the industry. However, knowledge about how to use it can be hard to come by, particularly if you’re looking for clear, simple, approachable terms. In this course, instructor Opeyemi Onikute provides a high-level overview of what Prometheus is and shows you how to install, use, scale, and manage it.</p><p>Discover the skills and core concepts you need to know to succeed in system monitoring with Prometheus. Explore the basics of setup, configuration, querying, defining rules, instrumentation, storage, and more. Upon completing this course, you’ll be prepared to continue your system monitoring journey with Prometheus.</p>


_See the readme file in the main branch for updated instructions and information._
## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Docker](https://docs.docker.com/desktop/install/mac-install/)
    - Any IDE of your choosing. e.g. [Visual Studio Code](https://visualstudio.microsoft.com/)
    - [Golang](https://go.dev/doc/install) - at least version 1.21
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.
3. When working with the sample HTTP server in `examples/http-server` (Chapter 2, Video 3 - Vector Matching), be sure to generate traffic using the simple bash script provided. This will ensure that metrics are populated when Prometheus scrapes the application:
    ```
    cd examples/http-server
    go run main.go
    ./generate-traffic.sh
    ```

### Instructor

Opeyemi Onikute

Systems Reliability Engineer at Cloudflare
                      

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/opeyemi-onikute?u=104).

[0]: # (Replace these placeholder URLs with actual course URLs)

[lil-course-url]: https://www.linkedin.com/learning/prometheus-essential-training
[lil-thumbnail-url]: https://media.licdn.com/dms/image/D4D0DAQE5GUkZ_WssLQ/learning-public-crop_675_1200/0/1712173826683?e=2147483647&v=beta&t=hAXP8NLRQTmtf3EduGg2GP2RvgUjQCbGUrQz5iYkk7k

