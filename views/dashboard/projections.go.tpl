{{ extends "page.go.tpl" }}

{{ define "title" }} Aperture - Personio {{ end }}

{{ define "content" }}

<div class="container-fluid">
    <div class="row">
        <div class="col main-nav bg-dark vh-100">
            <div class="dropdown w-100 d-block">
                <button aria-expanded="false"
                        class="py-2 btn btn-dark w-100 dropdown-toggle text-start d-flex align-items-center justify-content-between border-light border-bottom-2 border-top-0 border-start-0 border-end-0 rounded-0"
                        data-bs-toggle="dropdown" id="dropdownMenuButton2" type="button">
                    <span>
                    <i class="ti ti-box"></i>
                    &nbsp;Personio
                    </span>
                </button>
                <ul aria-labelledby="dropdownMenuButton2" class="dropdown-menu dropdown-menu-dark">
                    <li><a class="dropdown-item active" href="#">Action</a></li>
                    <li><a class="dropdown-item" href="#">Another action</a></li>
                    <li><a class="dropdown-item" href="#">Something else here</a></li>
                    <li>
                        <hr class="dropdown-divider">
                    </li>
                    <li><a class="dropdown-item" href="#">Separated link</a></li>
                </ul>
            </div>
            <ul class="nav flex-column mt-3 khand">
                <li class="nav-item">
                    <a aria-current="page" class="nav-link  text-light" href="#">

                        <i class="ti ti-layout-dashboard"></i>
                        &nbsp;Overview
                    </a>
                </li>
                <li class="nav-item">
                    <a aria-current="page" class="nav-link  text-light active" href="#">

                        <i class="ti ti-brand-youtube"></i>
                        &nbsp;Projections</a>
                </li>
                <li class="nav-item">
                    <a aria-current="page" class="nav-link  text-light" href="#">

                        <i class="ti ti-stack-2"></i>
                        &nbsp;Event Stream</a>
                </li>

            </ul>
        </div>
        <div class="col bg-light">
            <nav class="navbar navbar-expand-lg navbar-light bg-light">
                <div class="container-fluid px-0">
                    <!--                    <a class="navbar-brand" href="#">Navbar w/ text</a>-->
                    <!--                    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">-->
                    <!--                        <span class="navbar-toggler-icon"></span>-->
                    <!--                    </button>-->
                    <div class="collapse navbar-collapse khand khand-400" id="navbarText">
                        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                            <li class="nav-item">
                                <a aria-current="page" class="nav-link active" href="#">

                                    <i class="ti ti-brand-youtube"></i>
                                    &nbsp;Projections /</a>
                            </li>

                        </ul>
                        <span class="navbar-text">
                            v0.1.0
      </span>
                    </div>
                </div>
            </nav>
            <div class="container-fluid px-3 mt-1">
                <div class="row">
                    <div class="col ">
                        {{ range .Data.projections }}
                            {{ template "partial/projections/projection_item.go.tpl" . }}
                        {{ end }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

{{ end }}