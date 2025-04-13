<script lang="ts">
    import type { Domain } from '$lib/types/domain'

    export let data
    let domains: Domain[] = [...data.domains]
</script>

<div class="about-background">
    <div class="info-section" dir="rtl">
        <div class="collumns">
            <div class="collumn collumn-2">
                <h1 class="head headline hl1">عن المعجم الحويط</h1>
                <h3>ما هو المعجم الحويط؟</h3>
                <p>
                    معجم المصطلحات الحويط هو محاولة لتجميع المصطلحات المختلفة
                    التي تندرج تحت مجموعة معينة من المواضيع أو المجالات، وذلك
                    لتوفير مصدر شامل يتيح للمستخدمين الوصول بسهولة إلى معاني
                    المصطلحات وتفسيرها بشكل دقيق. هو أقرب إلى مصدر يجمع
                    المحاولات السابقة ويُلمّ الشتات الموجود في اللغة العربية في
                    المجال العلمي، ليكون مرجعًا موثوقًا يعكس تنظيم وتوحيد
                    المصطلحات في المجالات المتخصصة. كما يضم أيضًا استخدامات
                    المصطلحات كما تظهر في بعض الشركات الكبيرة مثل آبل وموزيلا،
                    التي ساهمت في تشكيل وتحديد شكل هذه المصطلحات. يهدف المعجم
                    إلى تسهيل البحث والفهم للمختصين والباحثين في تلك المجالات،
                    مما يسهم في تعزيز تبادل المعرفة والفهم المشترك بين المتحدثين
                    باللغة العربية.
                </p>
                <h3>اللغات المدعومة</h3>
                <div class="lang-list">
                    <span>العربية</span>
                    <span>الانجليزية</span>
                    <span>الفرنسية</span>
                    <span>الالمانية</span>
                </div>
                <p>
                    يحاول المعجم الحويط العثور على مصادر تحوي الأربع لغات التي
                    اعتبرها مربوطة أو مؤثرة على اللغة العربية في السياق العلمي،
                    ولكن في الأساس يشترط وجود الإنجليزية نظرًا لتأثيرها الحالي
                    على المجالات العلمية، خصوصًا البرمجية والنفسية. يعمل المعجم
                    على توثيق هذه اللغات من خلال قواميس متعددة اللغات والمراجع
                    العلمية التي تدمج بينها
                </p>
            </div>
            <div class="collumn collumn-3">
                <h3 class="text-3xl">المجالات</h3>
                <div>
                    <table style="margin-left: auto;margin-right: auto;">
                        <thead>
                            <tr>
                                <th scope="col" style="text-align: right;">
                                    عدد المفردات
                                </th>
                                <th scope="col" style="text-align: right;">
                                    العربية
                                </th>
                                <th scope="col" style="text-align: right;">
                                    الالمانية
                                </th>
                                <th scope="col" style="text-align: right;">
                                    الفرنسية
                                </th>
                                <th scope="col" style="text-align: right;">
                                    الانجليزية
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each domains as domain}
                                <tr>
                                    <td
                                        style="text-align: right;"
                                        data-th="عدد المفردات">{domain.TermCount}</td
                                    >
                                    <td
                                        style="text-align: right;"
                                        data-th="Arabic">{domain.NameAr}</td
                                    >
                                    <td class="al" data-th="German"
                                        >{domain.NameDE}</td
                                    >
                                    <td class="al" data-th="French"
                                        >{domain.NameFr}</td
                                    >
                                    <td class="al" data-th="English"
                                        >{domain.Name}</td
                                    >
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    /* tr {
        border-top: 1px solid #ddd;
        border-bottom: 1px solid #ddd;
        }
        
        th {
            display: none;
    }

    td {
        display: block;

        &:first-child {
            padding-top: 0.5em;
        }
        &:last-child {
            padding-bottom: 0.5em;
        }

        &:before {
            content: attr(data-th) ': ';
            font-weight: bold;

            width: 6.5em;
            display: inline-block;

            @media (min-width: 480px) {
                display: none;
            }
        }
    }

    th,
    td {
        text-align: left;

        @media (min-width: 480px) {
            display: table-cell;
            padding: 0.25em 0.5em;

            &:first-child {
                padding-left: 0;
            }

            &:last-child {
                padding-right: 0;
            }
        }
    } */
</style>
