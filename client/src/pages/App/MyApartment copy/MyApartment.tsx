import { Apartment } from "../../../DAL/services/apartments/apartments";

type Props = {
    apartment: Apartment | null
}

const MyApartment = (props: Props) => {
    const { apartment } = props
    return <h1>MyApartment</h1>;
}

export { MyApartment }