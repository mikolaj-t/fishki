<script>
    import Form from "../../../../../components/Form.svelte";
    import {page} from "$app/stores";
    import {Button, Input, Label, Textarea} from "flowbite-svelte";
    import {apiURL} from "../../../../../stuff.js";
    import CardEditor from "../../../../../components/CardEditor.svelte";
    import {beforeNavigate} from "$app/navigation";

    export let num = 1;

    /** @type {import('./$types').PageData} */
    export let data;

    async function FetchRequest(prompt, answer){
        const card = { prompt: prompt, answer: answer}
        const res = await fetch(apiURL + '/cards/create?deck=' + data.id, {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify(card)
        })

        window.location.href = "/deck/" + data.id
    }

    async function onSubmit(e) {
        const formData = new FormData(e.target);
        let i = 0;
        let prompt;
        let answer;
        for (const pair of formData.entries()) {
            if(i === 0) {
                prompt = pair[1];
            } else if(i === 1){
                answer = pair[1];
                FetchRequest(prompt, answer);
            }
            i = ++i % 2;
        }
    }

    function beforeUnload() {
        // Cancel the event as stated by the standard.
        event.preventDefault();
        // Chrome requires returnValue to be set.
        event.returnValue = '';
        // more compatibility
        return '...';
    }

    // I have no idea why, but it actually turns sveltekit's navigation change to a page unload,
    // so it calls beforeUnload() and hence the confirmation dialog shows up, that's nice
    beforeNavigate(() => { throw null; })
</script>


<svelte:window on:beforeunload={beforeUnload}/>
<Form>
    <form action="" on:submit|preventDefault={onSubmit}>
        {#each Array(num) as _, i}
            <CardEditor/>
        {/each}
        <Button class="addNext" on:click="{() => ++num}" gradient color="green">Add next</Button>
        <Button class="removeNext" on:click="{() => num > 1 ? --num : num = num}" gradient color="red">Remove next</Button>
        <Button type="submit" gradient color="blue">Create {num} {num === 1 ? "card" : "cards" }</Button>
    </form>
</Form>