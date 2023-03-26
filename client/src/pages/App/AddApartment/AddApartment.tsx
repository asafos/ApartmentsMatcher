import { Apartment } from "../../../DAL/services/auth/apartments";

type Props = {
    apartment: Apartment | null
}

const AddApartment = (props: Props) => {
    const { apartment } = props
    return <h1>AddApartment</h1>;
}

export { AddApartment }