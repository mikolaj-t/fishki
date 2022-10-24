<script>
import Form from "../../../components/Form.svelte";
import {Button, Input, Label} from "flowbite-svelte";
import {apiURL} from "../../../stuff.js";

async function onSubmit(e){
    const formData = new FormData(e.target);

    const res = await fetch(apiURL + '/user/register', {
        method: 'POST',
        body: formData,
    })

    if(res.ok){
        window.location.href = "/user/login";
    }
}

</script>

<Form>
    <form on:submit|preventDefault={onSubmit}>
        <Label for="username">Username</Label>
        <Input id="username" name="username" required/>
        <Label for="email">Email</Label>
        <Input id="email" name="email" type="email" required/>
        <Label for="password">Password</Label>
        <Input id="password" name="password" type="password" required/>
        <Button type="submit" gradient color="blue">Register</Button>
    </form>
</Form>