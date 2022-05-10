function timer () {
    const ListItem =  () =>{
        const number = [1, 2, 3, 4, 6]
        const list = number.map((number) =>
        <li>{number}</li>
        )
        return list
    }
    return (
        <div>
            <h2>Count: </h2>
             <ListItem/>
        </div>
    )
}
export default timer;