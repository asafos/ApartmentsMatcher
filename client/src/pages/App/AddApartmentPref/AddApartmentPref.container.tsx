import { ApartmentPrefToAdd } from '../../../DAL/services/apartmentPrefs/apartmentPrefs'
import { useApartmentPrefsStore } from '../../../DAL/stores/apartmentPrefs'
import { useAuthStore } from '../../../DAL/stores/auth'
import { AddApartmentPref } from './AddApartmentPref'

const AddApartmentPrefContainer = () => {
  const addUserApartmentPref = useApartmentPrefsStore(
    ({ addUserApartmentPref }) => addUserApartmentPref,
  )
  const user = useAuthStore(({ user }) => user)

  const handleSave = (apartmentPrefToAdd: Omit<ApartmentPrefToAdd, 'user_id'>) => {
    if (!user.data) {
      return
    }
    addUserApartmentPref({ ...apartmentPrefToAdd, user_id: user.data.id })
  }

  return <AddApartmentPref onSave={handleSave} />
}

export { AddApartmentPrefContainer }
