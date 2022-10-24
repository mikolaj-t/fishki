<script>
    import {onMount} from "svelte";
    import {differenceInDays} from "../routes/review/[id]/+page.js";
    import {apiURL} from "../stuff.js";

    export let id;
    export let review = {};
    export let todayAmount = 0

    async function f(){
        const res = await fetch(
            apiURL + '/reviews/get?id=' + id,
            {
                method: 'GET',
                credentials: 'include'
            }
        )
        const data = await res.json();
        console.log(data);

        if(res.ok){
            review = data;
        }

        let dateRef = new Date(2022, 1, 2);
        let currentDay = new Date(Date.now());
        let testDif = differenceInDays(dateRef, new Date(2022, 1, 3))
        let currentDayDif = differenceInDays(dateRef, currentDay);
        Object.entries(data.mode.dates).forEach((entry) => {
            if(entry[1] <= currentDayDif){
                todayAmount++;
            }
        });

        return review
    }
    let m = f()
    onMount(async () => {

    });
</script>

{#await m}
    {:then m}
    <a href="/review/{review.id}"> {review.name} - <a class="text-green-400 font-extrabold">{todayAmount}</a> today!</a> <br/>
    {/await}
