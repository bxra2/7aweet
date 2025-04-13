<script>
    import '../app.css'
</script>

<div class="head">
    <div dir="ltr" class="headerobjectswrapper">
        <div class="weatherforcastbox">
            <span> معجم هدفه لم شتات المصطلحات في المجالات العلمية </span>
        </div>
        <header>المعجم الحويط</header>
    </div>

    <div class="subhead">
        <nav>
            <ul>
                <li>
                    <a href="/">الرئيسية</a>
                </li>
                <li>
                    <a href="/sources">المصادر</a>
                </li>
                <li>
                    <a href="/suggestions">المقترحات</a>
                </li>
                <li>
                    <a href="/about">من نحن</a>
                </li>
            </ul>
        </nav>
    </div>
</div>
<div style="height: 100%;">
    <slot></slot>
</div>
<!-- 
<style>
    :root {
        --background-color: #f9f7f1;
        --color: #2f2f2f;
    }

    /* #9d9487
#473625 */

    * {
        font-family: 'Playfair Display', serif !important;
    }
    body {
        font-family: 'Droid Serif', serif;
        font-size: 14px;
        color: var(--color);
        background-color: var(background-color);
    }

    header {
        font-family: 'Playfair Display', serif;
        font-weight: 900;
        font-size: 80px;
        text-transform: uppercase;
        display: inline-block;
        line-height: 72px;
        margin-bottom: 20px;
    }

    p {
        margin-top: 0;
        margin-bottom: 20px;
    }

    .head {
        text-align: center;
        position: relative;
    }

    /* .headerobjectswrapper {
    } */

    .subhead {
        text-transform: uppercase;
        border-bottom: 3px double var(--color);
        border-top: 3px double var(--color);
        padding: 12px 0 12px 0;
    }

    /* .subhead nav li:not(:last-child)::after {
        content: '';
        position: absolute;
        right: 70px;
        top: 25%;
        width: 3px;
        height: 50%;
        background-color: var(--color, #ccc);
    } */

    nav {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    nav ul {
        display: flex;
        justify-content: center;
        list-style: none;
        padding: 0;
        margin: 0;
    }

    nav ul li {
        margin: 0 35px;
        font-size: 1.2em;
        position: relative;
    }

    nav a {
        text-decoration: none;
        color: inherit;
    }

    .weatherforcastbox {
        position: relative;
        width: 12%;
        left: 10px;
        border: 3px double var(--color);
        padding: 10px 15px 10px 15px;
        line-height: 20px;
        display: inline-block;
        margin: 0 50px 20px -360px;
    }
    .content {
        font-size: 0;
        line-height: 0;
        word-spacing: -0.31em;
        display: inline-block;
        margin: 30px 2% 0 2%;
    }
    .collumns {
    }

    .collumn {
        font-size: 14px;
        line-height: 20px;
        width: 17.5%;
        display: inline-block;
        padding: 0 1% 0 1%;
        vertical-align: top;
        margin-bottom: 50px;
        transition: all 0.7s;
    }
    .collumn + .collumn {
        border-left: 1px solid var(--color);
    }
    .collumn .headline {
        text-align: center;
        line-height: normal;
        font-family: 'Playfair Display', serif;
        display: block;
        margin: 0 auto;
    }
    .collumn .headline.hl1 {
        font-weight: 700;
        font-size: 30px;
        text-transform: uppercase;
        padding: 10px 0 10px 0;
    }

    .collumn .headline.hl2 {
        font-weight: 400;
        font-style: italic;
        font-size: 24px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }

    .collumn .headline.hl2:before {
        border-top: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 7px;
        display: block;
        margin: 0 auto;
    }
    .collumn .headline.hl2:after {
        border-bottom: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 13px;
        display: block;
        margin: 0 auto;
    }

    .collumn .headline.hl3 {
        font-weight: 400;
        font-style: italic;
        font-size: 36px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl4 {
        font-weight: 700;
        font-size: 12px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl4:before {
        border-top: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 7px;
        display: block;
        margin: 0 auto;
    }
    .collumn .headline.hl4:after {
        border-bottom: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 10px;
        display: block;
        margin: 0 auto;
    }

    .collumn .headline.hl5 {
        font-weight: 400;
        font-size: 42px;
        text-transform: uppercase;
        font-style: italic;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl6 {
        font-weight: 400;
        font-size: 18px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl6:before {
        border-top: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 7px;
        display: block;
        margin: 0 auto;
    }
    .collumn .headline.hl6:after {
        border-bottom: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 10px;
        display: block;
        margin: 0 auto;
    }
    .collumn .headline.hl7 {
        font-weight: 700;
        font-size: 12px;
        box-sizing: border-box;
        display: block;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl8 {
        font-weight: 700;
        font-size: 12px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl9 {
        font-weight: 700;
        font-size: 12px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .headline.hl10 {
        font-weight: 700;
        font-size: 12px;
        box-sizing: border-box;
        padding: 10px 0 10px 0;
    }
    .collumn .citation {
        font-family: 'Playfair Display', serif;
        font-size: 36px;
        line-height: 44px;
        text-align: center;
        font-weight: 400;
        display: block;
        margin: 40px 0 40px 0;
        font-feature-settings: 'liga', 'dlig';
    }
    .collumn .citation:before {
        border-top: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 16px;
        display: block;
        margin: 0 auto;
    }
    .collumn .citation:after {
        border-bottom: 1px solid var(--color);
        content: '';
        width: 100px;
        height: 16px;
        display: block;
        margin: 0 auto;
    }
    .collumn .figure {
        margin: 0 0 20px;
    }
    .collumn .figcaption {
        font-style: italic;
        font-size: 12px;
    }
    .media {
        -webkit-filter: sepia(80%) contrast(1) opacity(0.8);
        filter: sepia(80%) grayscale(1) contrast(1) opacity(0.8);
        mix-blend-mode: multiply;
        width: 100%;
    }
    @media only all and (max-width: 1300px) {
        .weatherforcastbox {
            display: none;
        }
    }
    @media only all and (max-width: 1200px) {
        .collumn {
            width: 31%;
        }
    }
    @media only all and (max-width: 900px) {
        .collumn {
            width: 47%;
        }
    }
    @media only all and (max-width: 600px) {
        .collumn {
            width: 100%;
        }
        .collumn + .collumn {
            border-left: none;
            border-bottom: 1px solid var(--color);
        }
        header {
            max-width: 320px;
            font-size: 60px;
            line-height: 54px;
            overflow: hidden;
        }
    }
</style> -->
