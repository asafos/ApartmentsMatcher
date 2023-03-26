import { useEffect } from "react";
import { Apartment } from "../../../DAL/services/auth/apartments";
import { useApartmentsStore } from "../../../DAL/stores/apartments";
import { useAuthStore } from "../../../DAL/stores/auth";
import { AddApartment } from "./AddApartment";

type Props = {
}

const AddApartmentContainer = (props: Props) => {
    const { } = props
    const apartment = useApartmentsStore(({ apartment }) => apartment)

    return <AddApartment apartment={apartment.data} />;
}

export { AddApartmentContainer }