<script>
import Form from "../../../components/Form.svelte";
import {Button, Input, Label} from "flowbite-svelte";
import {apiURL} from "../../../stuff.js";

async function Submit(e){
    const formData = new FormData(e.target);

    const data = {
        name: formData.get("name"),
        'public': true,
    }

    const res = await fetch(apiURL + '/decks/create', {
        method: 'POST',
        body: JSON.stringify(data),
        credentials: 'include'
    })

    const resData = await res.json();
    console.log(resData);
    window.location.href = "/deck/" + resData.id;

    //("alert");
    if(res.ok){
    }
}
</script>

<Form>
    <form on:submit|preventDefault={Submit} class="columns-1">
        <Label for="name" >Name a new deck.,.</Label>
        <Input name="name" id="name" placeholder="ex. guinea pigs"/>
        <Button class="mt-3" type="submit" gradient color="blue">Create</Button>
    </form>
</Form>