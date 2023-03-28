import { ApartmentToAdd } from '../../../DAL/services/apartments/apartments'
import { useApartmentsStore } from '../../../DAL/stores/apartments'
import { useAuthStore } from '../../../DAL/stores/auth'
import { AddApartment } from './AddApartment'

const AddApartmentContainer = () => {
  const addUserApartment = useApartmentsStore(({ addUserApartment }) => addUserApartment)
  const user = useAuthStore(({ user }) => user)

  const handleSave = (apartmentToAdd: Omit<ApartmentToAdd, 'user_id'>) => {
    if (!user.data) {
      return
    }
    addUserApartment({ ...apartmentToAdd, user_id: user.data.id })
  }

  return <AddApartment onSave={handleSave} />
}

export { AddApartmentContainer }
