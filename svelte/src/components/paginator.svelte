<script lang="ts">
    export let currentPage: number = 1
    export let totalPages: number = 10 // Can be passed from parent or calculated
    export let query: string = '' // Your search query
    export let onPageChange: (page: number) => void // Callback function to handle page change

    const pagesToShow = 3 // Show 3 pages before and after the current page

    // Calculate the page range to display
    $: pageRange = calculatePageRange(currentPage, totalPages, pagesToShow)

    // Function to calculate which pages to display
    function calculatePageRange(current: number, total: number, range: number) {
        let start = Math.max(1, current - range) // Ensure we don't go below 1
        let end = Math.min(total, current + range) // Ensure we don't go above total pages

        let pages = []
        for (let i = start; i <= end; i++) {
            pages.push(i)
        }

        return pages
    }

    // Function to handle page click (this triggers the onPageChange function)
    function goToPage(page: number) {
        if (page >= 1 && page <= totalPages) {
            onPageChange(page)
            currentPage = page
        }
    }

    // Handle "Next" and "Previous" button logic
    function nextPage() {
        if (currentPage < totalPages) {
            goToPage(currentPage + 1)
        }
    }

    function prevPage() {
        if (currentPage > 1) {
            goToPage(currentPage - 1)
        }
    }
</script>

<div class="pagination">
    <!-- Previous Button -->
    <a
        href="/search?q={query}&page={currentPage - 1}"
        class:disabled={currentPage === 1}
        on:click={(e) => {
          if (currentPage === 1) {
              e.preventDefault()
          }
      }}
    >
        السابق
    </a>

    <!-- Page numbers -->
    {#each pageRange as page}
        <a
            class={page === currentPage ? 'active' : ''}
            href="/search?q={query}&page={page}"
        >
            {page}
        </a>
    {/each}

    <!-- Next Button -->
    <a
        href="/search?q={query}&page={currentPage + 1}"
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

<style>
    .pagination {
        display: flex;
        justify-content: center;
        gap: 10px;
        align-items: center;
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

    a.active {
        background-color: #4a4a4a;
        color: white;
    }

    a:hover {
        background-color: #eeeeee;
    }

    a:focus {
        outline: none;
    }
</style>
