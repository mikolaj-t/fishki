/** @type {import('./$types').PageLoad} */
export async function load({params}){
    return {
        id: params.id
    }
}