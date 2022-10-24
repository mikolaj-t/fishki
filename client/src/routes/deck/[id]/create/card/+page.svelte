<script>
import Form from "../../../../../components/Form.svelte";
import {page} from "$app/stores";
import {Button, Input, Label} from "flowbite-svelte";
import {apiURL} from "../../../../../stuff.js";

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

function onSubmit(e) {
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
</script>

<Form>
    <form action=""c on:submit|preventDefault={onSubmit}>
    {#each Array(num) as _, i}
        <div class="grid gap-6 mb-6 md:grid-cols-2">
            <div>
            <Label for="prompt">🐟 Prompt</Label>
            <Input id="prompt" name="prompt" placeholder="Coolest animal on earth" required/>
            </div>
            <div>
            <Label for="answer">🗝️ Answer</Label>
            <Input id="answer" name="answer" placeholder="turtle" n required/>
            </div>
        </div>
    {/each}
        <Button class="addNext" on:click="{() => ++num}" gradient color="green">Add next</Button>
        <Button class="removeNext" on:click="{() => num > 1 ? --num : num = num}" gradient color="red">Remove next</Button>
        <Button type="submit" gradient color="blue">Create {num} {num === 1 ? "card" : "cards" }</Button>
    </form>
 </Form>
