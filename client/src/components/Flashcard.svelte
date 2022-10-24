<script context="module">
    export function LetsTryThis(){
            alert("aa");
    }
</script>
<script>

    import { onMount } from 'svelte';
    import Sticky from "./Sticky.svelte";
    import {currentCardID} from "../stores.js";
    import {apiURL} from "../stuff.js";

    export let flashcard = { prompt: "", answer:""};
    export let id;
    export let sticky = false;

    export let nextCardFunc;

    let toggle = false
    const toggleCard = () => toggle = !toggle;

    // $: alert(id);

    onMount(async () => {
        const res = await fetch(
            apiURL + '/cards/get?id=' + id
        )
        const data = await res.json();
        console.log(data);

        if(res.ok){
            flashcard = data;
        }
    })

    $: fetch(apiURL + '/cards/get?id=' + id)
        .then(response => response.json())
        .then(data => {
            flashcard.prompt = data.prompt;
            flashcard.answer = data.answer;
        })
</script>

<div class="scene centered">
    <div class="flashcard" class:flip={toggle} on:click="{toggleCard}">
        <div class="face front">
            <div class="top">🐟</div>
            <div class="text">{flashcard.prompt}</div>
        </div>
        <div class="face back">
            <slot></slot>
            {#if sticky}
                <Sticky nextCardFunc="{nextCardFunc}"></Sticky>
            {/if}
            <div class="top">🗝️</div>
            <div class="text" class:hidden={!toggle} on:click="{toggleCard}">{flashcard.answer}</div>
        </div>
    </div>
</div>

<style>
    .hidden {
        visibility: hidden;
    }
    .scene {
        padding-bottom: 0vh;
        perspective: 5000px;
        height: 90vh;
        aspect-ratio: 1 / 1.414213562373095048801688724209;
        max-width: 100vw;
    }

    .top {
        position: absolute;
        top: 0;
        margin-top: 3vh;
        font-size: 6vh;
    }
    .flashcard {
        width: 100%;
        height: 100%;
        box-shadow: 0px 0px 3px 1px #d7d7d7;
        border-radius: 5px 5px 5px 5px;
        position: relative;
        transition-duration: 0.5s;
        transition-property: transform;
        transform-style: preserve-3d;
        text-align: center;
    }

    .face {
        position: absolute;
        height: 100%;
        width: 100%;
        backface-visibility: hidden;
        background: #ffffff;

        border-radius: 5px 5px 5px 5px;

        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 30pt;

    }

    .text {
        padding: 50px;
        font-size: 3.5vh;
        text-align: center;
    }

    .front {
        font-weight: 600;
        /*color: #0050fd;*/
    }

    .back{
        transform: rotateY(180deg);
        /*color: goldenrod;*/
    }

    .flashcard.flip {
        transform: rotateY(180deg);
    }

    /*.flashcard:hover {
        transform: scale(102%);
    }*/
</style>