<script>
    export let inputs;
    export let includeForm = false;
</script>

{#if includeForm}
<form>
    <slot></slot>
</form>
{:else }
    <slot></slot>
{/if}

<style>
    form {
        display: grid;
        grid-gap: 1vh;
        width: max-content;
    }
    .inline {
        display: inline-grid;
    }
    form :global(input) {
        padding: 0.5rem
    }

    form :global(select) {
        padding: 0.5rem
    }


</style>