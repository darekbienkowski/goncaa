<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">goncaa</h3>

  <p align="center">
    A terminal based application for viewing live and past NCAA games and statistics
    <br />
    <br />
    <br /> 
    <a href="https://github.com/darekbienkowski/goncaa/issues">Report Bug</a>
    ·
    <a href="https://github.com/darekbienkowski/goncaa/issues">Request Feature</a>

[![GPL License][license-shield]][license-url]

  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
        <a href="#installation">Installation</a>
    </li>
    <li><a href="#usage">Usage</a></li> 
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

![goncaa Screen Shot][product-screenshot]

`goncaa` is a terminal based user interface that allows you to scroll through both live and completed NCAA games and browse scores, players, and statistics.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

- Go
- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Bubble-table](https://github.com/Evertras/bubble-table)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Installation

```bash
go install github.com/darekbienkowski/goncaa@latest
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

By default running the application will open the game list view with the games from today's date. If you want to specify a specific date to open up to you can do so by doing

```bash
goncaa
goncaa -d YYYY-MM-DD
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the GPL License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments
### Original MLB Project

This project is a fork of [AxBolduc/gomlb][gomlb_url].    
While the sport, models, and API sources have been updated, the core logic remains largely based on the original codebase.

- Project Link: [https://github.com/AxBolduc/gomlb](https://github.com/AxBolduc/gomlb)
---
- https://github.com/henrygd/ncaa-api fantastic NCAA API used by goncaa  
- [nbacli](https://github.com/dylantientcheu/nbacli/) inspiration for gomlb and code snippets

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[gomlb_url]: https://github.com/AxBolduc/gomlb
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://github.com/darekbienkowski/goncaa/blob/main/LICENSE
[product-screenshot]: images/repo_image.png
[product-demo]: images/repo_gif.gif
