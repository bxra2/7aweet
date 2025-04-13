<script>
    import SearchForm from '../components/searchForm.svelte'
</script>

<div class=" flex-center">
    <SearchForm />
</div>

<form class="newspaper-form">
    <input type="search" placeholder="Search..." class="search-input" />

    <div class="checkbox-group">
        <label><input type="checkbox" value="1" /> وصف</label>
        <label><input type="checkbox" value="1" /> فرنسي</label>
        <label><input type="checkbox" value="1" /> الماني</label>
    </div>

    <button type="submit" class="newspaper-button">ابحث</button>
</form>

<style>
    .newspaper-form {
        font-family: 'Times New Roman', Times, serif;
        background-color: #fff;
        border: 2px solid #000;
        padding: 20px;
        max-width: 400px;
        box-shadow: 5px 5px 0 #000;
        display: flex;
        flex-direction: column;
        gap: 15px;
    }

    .search-input {
        padding: 10px;
        font-size: 16px;
        border: 2px solid #000;
        background-color: #fefefe;
        color: #000;
        font-family: inherit;
    }

    .checkbox-group {
        display: flex;
        flex-direction: column;
        gap: 5px;
        font-size: 14px;
        text-transform: uppercase;
        letter-spacing: 0.5px;
    }

    .checkbox-group input[type='checkbox'] {
        margin-right: 8px;
        accent-color: black; /* modern browsers support this */
    }

    .newspaper-button {
        font-family: 'Times New Roman', Times, serif;
        font-size: 18px;
        font-weight: bold;
        color: #000;
        background-color: #fff;
        border: 2px solid #000;
        padding: 10px 20px;
        text-transform: uppercase;
        letter-spacing: 1px;
        box-shadow: 2px 2px 0 #000;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    .newspaper-button:hover {
        background-color: #f5f5f5;
        box-shadow: none;
        transform: translate(2px, 2px);
    }
</style>
