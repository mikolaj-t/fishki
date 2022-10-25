<script>
    import CardEditor from "../../../../components/CardEditor.svelte";
    import {apiURL} from "../../../../stuff.js";
    import {Button} from "flowbite-svelte";

    /** @type {import('./$types').PageData} */
    export let data;

    async function onSubmit(e){
        let formData = new FormData(e.target);
        FetchRequest(formData.get("prompt"), formData.get("answer"));
    }

    async function FetchRequest(prompt, answer){
        const card = { id:data.id,prompt: prompt, answer: answer}
        const res = await fetch(apiURL + '/cards/update', {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify(card)
        })
        console.log(res)
        window.location.href = "/deck/" + data.id
    }
</script>

<form action="" on:submit|preventDefault={onSubmit}>
    <CardEditor id={data.id}/>
    <Button type="submit" gradient color="blue">Edit card</Button>

</form>