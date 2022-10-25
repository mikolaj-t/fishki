<script>
    import {Label, Textarea} from "flowbite-svelte";
    import {onMount} from "svelte";
    import {apiURL} from "../stuff.js";

    export let id;
    export let flashcard = {prompt: undefined, answer:undefined};

    onMount(async () => {
        console.log("card editor");
        const res = await fetch(
            apiURL + '/cards/get?id=' + id
        )
        const data = await res.json();
        console.log(data);

        if(res.ok){
            flashcard = data;
        }
    })

</script>

<div class="grid gap-6 mb-6 md:grid-cols-2">
    <div>
        <Label for="prompt">🐟 Prompt</Label>
        <Textarea value={flashcard.prompt === undefined ? '' : flashcard.prompt} id="prompt" name="prompt" placeholder="Coolest animal on earth" required/>
        <!--{flashcard.prompt === undefined ? '' : flashcard.prompt}-->
    </div>
    <div>
        <Label for="answer">🗝️ Answer</Label>
        <Textarea value={flashcard.answer === undefined ? '' : flashcard.answer} id="answer" name="answer" placeholder="turtle" n required/>
    </div>
</div>