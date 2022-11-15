<script>
    import Form from "../../../../../components/Form.svelte";
    import {redirect} from "@sveltejs/kit";
    import {Button, Input, Label, Select} from "flowbite-svelte";
    import {apiURL} from "../../../../../stuff.js";

    /*let modes = [
        {id:0, text: ''},
        { id: 1, text: `Fixed` },
    ];*/

    let modes = [
        {value: 1, name: "fixed"}
    ]

    let selected;

    /** @type {import('./$types').PageData} */
    export let data;


    let answer = '';

    //modeID deckID mode.intervals
    async function onSubmit(e){
        const formData = new FormData(e.target);
        /*const body = { modeID: 1, deckID: data.id, mode: {
            intervals: formData.get("mode.intervals").split(",").map(Number)
        }};*/
        formData.set("modeID", "1");
        formData.set("deckID", data.id);
        formData.get("mode.intervals").split(",").forEach(value => {formData.append("intervals", value)});
        formData.delete("mode.intervals");
        console.log(formData);

        const res = await fetch(apiURL + '/reviews/create', {
            method: 'POST',
            credentials: 'include',
            body: formData
        })

        const ddd = await res.json();
        let review;
        if(res.ok) {
            review = ddd;
        }

        window.location.href = "/review/" + review.id;
    }
</script>

<Form>
    <form on:submit|preventDefault={onSubmit}>
        <Label for="name">Review name</Label>
        <Input id="name" name="name" placeholder="a review of this awesome subject for this awesome test"/>
    <Label for="mode">Choose mode..</Label>
    <Select id="mode" items={modes} bind:value={selected} on:change="{() => answer = ''}" required>
        <!--{#each modes as mode}
            <option value={mode}>
                {mode.text}
            </option>
        {/each}-->
    </Select>
    <!--<select id="mode">
        <option>fixed</option>
    </select>-->

    {#if selected && selected === 1}
        <Label for="intervals"> Intervals </Label>
        <Input id="intervals" name="mode.intervals" placeholder="1,2,3" required/>
        <Button type="submit" gradient color="blue">Create review</Button>
    {/if}
    </form>
</Form>