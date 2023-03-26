import { useEffect } from "react";
import { Apartment } from "../../../DAL/services/apartments/apartments";
import { useApartmentsStore } from "../../../DAL/stores/apartments";
import { useAuthStore } from "../../../DAL/stores/auth";
import { MyApartment } from "./MyApartment";

type Props = {
}

const MyApartmentContainer = (props: Props) => {
    const { } = props
    const apartment = useApartmentsStore(({ apartment }) => apartment)

    return <MyApartment apartment={apartment.data} />;
}

export { MyApartmentContainer }