<script>
    import Fa from 'svelte-fa/src/fa.svelte'
    import { faCircleCheck } from '@fortawesome/free-solid-svg-icons';
    import { faCircleXmark } from '@fortawesome/free-solid-svg-icons';
    import {viewerCardID} from "../stores.js";
    import {SubmitAnswer} from "../stuff.js";

    function CorrectAnswer(){
        SubmitAnswer(true);
        NewCardFunc();
    }

    function IncorrectAnswer(){
        SubmitAnswer(false);
        NewCardFunc();
    }
    function NewCardFunc(){
        viewerCardID.update(n => n + 1);
    }
</script>

<div class="sticky">
    <button class="success btn" on:click="{() => CorrectAnswer()}">
        <Fa icon="{faCircleCheck}" primaryColor="#9ccc65" size="10vh"/>
        <!--<div class="text">Knew it</div-->
    </button>
    <button class="fail btn" on:click="{() => IncorrectAnswer()}">
        <Fa icon="{faCircleXmark}" primaryColor="#ef5350" size="10vh"/>
        <!--<div class="text">Forgot it</div>-->
    </button>
</div>

<style>
    .sticky {
        bottom: 0;
        position: absolute;
        justify-content: center;
        display: flex;
        width: 100%;
        height: 15vh;
        max-height: 200px;
        padding-left: 2%;
        padding-right: 2%;
        margin-bottom: 2%;
    }

    .btn {
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 100%;
        margin: 0 5%;
        max-width: 200px;
        width: min-content;
        height: min-content;
        aspect-ratio: 1/1;
        border: none;
        background-color: #f5f5f5
        /*background-color: rgba(0, 0, 0, 0);*/
    }
    .success {
        /*width: 50%;*/
        /*background-color:#9ccc65;*/
    }
    .fail {
        /*width: 50%;*/
        /*background-color: #ef5350;*/
    }
    .text {
        font-size: 3vh;
        font-weight: 500;
        color: white;
        /*text-shadow: 0px 0px 2px #333333;*/
    }
</style>