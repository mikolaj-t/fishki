<script>
	import '../app.postcss';
    import {Navbar, NavBrand, NavHamburger, NavLi, NavUl} from "flowbite-svelte";
    import {Container} from "postcss";
</script>

<Navbar let:hidden let:toggle>
    <NavBrand href="/">
        <img
                src="https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/240/apple/325/fish_1f41f.png"
                class="mr-3 h-6 sm:h-9"
                alt="fishki"
        />
        <span class="self-center whitespace-nowrap text-xl font-semibold text-cyan-900 dark:text-white">
      fishki
        </span>
    </NavBrand>
    <NavHamburger on:click={toggle} />
    <NavUl {hidden}>
        <NavLi href="/" active={true}>Home</NavLi>
        <NavLi href="/me">Me</NavLi>
        <NavLi href="/user/login">Login</NavLi>
        <NavLi href="/user/register">Register</NavLi>
        <NavLi href="/finek">finek</NavLi>

    </NavUl>
</Navbar>

<main class="container mx-auto">
<slot/>
</main>
