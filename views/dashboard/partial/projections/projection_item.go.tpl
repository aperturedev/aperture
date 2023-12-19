<div class="d-flex flex-column flex-lg-row align-items-center justify-content-between my-3 projection-item">
    <div>
        <h3>{{ .Name }}</h3>
        <p>{{ .Description }}</p>
    </div>
    <div class="d-flex justify-content-end align-items-center pb-3 projection-actions">
        {{ if eq .State 5 }}
        <button class="btn btn-light khand-500" type="button">
            <i class="ti ti-circle-dot-filled text-success"></i>
            &nbsp;Running
        </button>
        {{ else }}
        <button class="btn btn-light khand-500" type="button">
            <i class="ti ti-alert-octagon-filled text-danger"></i>
            &nbsp;Failing
        </button>
        {{ end }}

        <span class="khand-500 projection-progress">
            <span class="text-danger">34</span> / <span class="text-success">3700</span>
        </span>
        <button class="btn btn-light btn-icon-only" type="button">
            <i class="ti ti-code text-dark"></i>
        </button>
        <button class="btn btn-warning" type="button">
            <i class="ti ti-clock-pause"></i>
            &nbsp;Pause
        </button>
        <button class="btn btn-danger" type="button">
            <i class="ti ti-refresh-dot"></i>
            &nbsp;Replay
        </button>
    </div>
</div>