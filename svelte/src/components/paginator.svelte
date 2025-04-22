<script lang="ts">
    import { toArabicIndic } from '$lib/utils/numerals'
    export let currentPage: number = 1
    export let limit: number = 10
    export let count: number
    export let totalPages: number = 10 // Can be passed from parent or calculated
    export let query: string = '' // Your search query

    const pagesToShow = 2 // Show 3 pages before and after the current page

    // Calculate the page range to display
    $: pageRange = calculatePageRange(currentPage, totalPages, pagesToShow)
    function calculatePageRange(current: number, total: number, range: number) {
        let start = current - range
        let end = current + range

        // Shift the range if start goes below 1
        if (start < 1) {
            end += 1 - start
            start = 1
        }

        // Shift the range if end goes beyond total
        if (end > total) {
            start -= end - total
            end = total
        }

        // Ensure start isn't less than 1 after shifting
        start = Math.max(1, start)

        const pages = []
        for (let i = start; i <= end; i++) {
            pages.push(i)
        }

        return pages
    }

    $: totalPages = Math.ceil(count / +limit)
    const limitsList = [5, 10, 20, 50]
</script>

<div class="pagination">
    <div class="pages">
        <!--  الصفحة السابقة  -->
        <a
            href="/search?q={query}&page={currentPage - 1}&limit={limit}"
            class:disabled={currentPage === 1}
            on:click={(e) => {
                if (currentPage === 1) {
                    e.preventDefault()
                }
            }}
        >
            السابق
        </a>
        <!--  الصفحة الاولي  -->
        {#if currentPage > 4}
            <a
                href="/search?q={query}&page={1}&limit={limit}"
                class:disabled={currentPage === 1}
                on:click={(e) => {
                    if (currentPage === 1) {
                        e.preventDefault()
                    }
                }}
            >
                {toArabicIndic(1)}
            </a>
            <span>...</span>
        {/if}

        <!--  الصفحات  -->
        {#each pageRange as page}
            <a
                class={page === currentPage ? 'active' : ''}
                class:current-disabled={currentPage === page}
                href="/search?q={query}&page={page}&limit={limit}"
                on:click={(e) => {
                    if (page === currentPage) {
                        e.preventDefault()
                    }
                }}
            >
                {toArabicIndic(page)}
            </a>
        {/each}

        <!--  الصفحة الاخيرة  -->
        {#if currentPage + 4 < totalPages}
            <span>...</span>
            <a
                href="/search?q={query}&page={totalPages}&limit={limit}"
                class:disabled={currentPage === totalPages}
                on:click={(e) => {
                    if (currentPage === totalPages) {
                        e.preventDefault()
                    }
                }}
            >
                {toArabicIndic(+totalPages)}
            </a>
        {/if}

        <!--  الصفحة التالية  -->
        <a
            href="/search?q={query}&page={currentPage + 1}&limit={limit}"
            class:disabled={currentPage === totalPages}
            on:click={(e) => {
                if (currentPage === totalPages) {
                    e.preventDefault()
                }
            }}
        >
            التالي
        </a>
    </div>
    <div class="limit">
        {#if limitsList[0] < count}
            <span>عدد النتائج</span>
        {/if}

        {#each limitsList as limitItem}
            {#if limitItem < count}
                <a
                    href="/search?q={query}&page={1}&limit={limitItem}"
                    class={limit === limitItem ? 'active' : ''}
                    class:current-disabled={limit === limitItem}
                    on:click={(e) => {
                        if (limit === limitItem) {
                            e.preventDefault()
                        }
                    }}
                >
                    {toArabicIndic(limitItem)}
                </a>
            {/if}
        {/each}
    </div>
</div>

<style>
    .pagination {
        display: flex;
        justify-content: space-around;
        align-items: center;
        margin-top: 20px;
    }
    .pages {
        display: flex;
        justify-content: center;
        gap: 10px;
        align-items: center;
    }

    .limit {
        gap: 10px;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    a {
        padding: 6px 8px;
        text-decoration: none;
        border: 1px solid #ccc;
        border-radius: 4px;
        cursor: pointer;
        color: black;
        box-shadow: 3px 3px 0 var(--color);
    }

    a.disabled {
        cursor: not-allowed;
        opacity: 0.5;
        pointer-events: none;
    }
    a.current-disabled {
        cursor: not-allowed;
        pointer-events: none;
    }

    a.active {
        background-color: #4a4a4a;
        color: white;
    }
    a.active:hover {
        background-color: #6e6e6e;
    }
    a:hover {
        background-color: #eeeeee;
    }

    a:focus {
        outline: none;
    }
</style>
