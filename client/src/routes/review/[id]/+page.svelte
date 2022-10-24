<script>
    import CardViewer from "../../../components/CardViewer.svelte";

    /** @type {import('./$types').PageData} */
    export let data;

</script>

<CardViewer IDs={data.cards}/>