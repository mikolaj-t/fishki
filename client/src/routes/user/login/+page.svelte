<script>
import Form from "../../../components/Form.svelte";
import {Button, Input, Label} from "flowbite-svelte";
import {apiURL} from "../../../stuff.js";

async function onSubmit(e){
    const formData = new FormData(e.target);

    const res = await fetch(apiURL + '/user/login', {
        method: 'POST',
        body: formData,
        credentials: 'include'
    })

    if(res.ok){
        window.localStorage.setItem("username", formData.get("username"));
        window.location.href = "/me";
    }
}

</script>

<Form>
    <form on:submit|preventDefault={onSubmit}>
        <Label for="username">Username</Label>
        <Input id="username" name="username" required/>
        <Label for="password">Password</Label>
        <Input id="password" name="password" type="password" required/>
        <Button type="submit"  gradient color="blue">Login</Button>
    </form>
</Form>