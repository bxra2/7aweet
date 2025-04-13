<!-- partial -->
<script>
    export let french, arabic, german, english, desc, sourceName, DomainName
</script>

<div class="main-content">
    <div class="ticket">
        <div class="header">
            <span class="ar">
                {arabic}
            </span>
            <span class="al">
                {english}
            </span>
        </div>
        <div class="info passenger">
            <div class="info__item">وصف</div>
            <div class="info__detail">{desc}</div>
        </div>
        <div class="info platform">
            <span>المصدر:</span>
            <span>{sourceName}</span>
            <hr />
            <span>المجال:</span>
            <span>{DomainName}</span>
        </div>
        <div dir="ltr" class="info departure">
            <div class="info__item">German</div>
            <div class="info__detail">{german}</div>
        </div>
        <div dir="ltr" class="info arrival">
            <div class="info__item">French</div>
            <div class="info__detail">{french}</div>
        </div>
    </div>
</div>

<style>
    .main-content {
        margin: 4em auto 0;
        width: 740px;
        text-transform: uppercase;
    }

    .ticket {
        display: grid;
        grid-template-columns: auto 143px;
        background: url(https://s3-us-west-2.amazonaws.com/s.cdpn.io/567707/paper_fibers.png)
            repeat;
        border-radius: 10px;
        border: 2px solid #222;
        cursor: default;
        padding: 10px;
    }
    .header {
        grid-area: title;
        grid-column: span 7;
        grid-row: 1;
        display: flex;
        justify-content: space-between;
        padding: 10px 20px;
        background: #04110f;
        color: #eee;
    }

    .info {
        border: 3px solid;
        border-width: 0 3px 3px 0;
        padding: 8px;
    }
    .info__item {
        font:
            400 13px 'Questrial',
            sans-serif;
        letter-spacing: 0.5px;
    }
    .info__detail {
        /* font: 700 20px/1 'Jura'; */
        font-weight: 500;
        margin: 4px 0;
    }

    .passenger {
        grid-column: 1 / span 6;
    }

    .platform {
        grid-column: 7 / span 1;
        grid-row: 2 / span 3;
        background: #1c1f33;
        color: #eee;
        border-color: #000;
        text-align: center;
        padding: 10px 0;
    }
    .platform span {
        display: block;
        font-size: 0.8rem;
        text-align: right;
        padding: 0 10px;
    }

    .departure,
    .arrival {
        grid-column-start: span 3;
    }

    .passenger,
    .departure {
        border-left: 3px solid;
    }
</style>
