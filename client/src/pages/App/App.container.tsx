import { useEffect } from 'react'
import { useNavigate, useLocation } from 'react-router-dom'
import { useApartmentPrefsStore } from '../../DAL/stores/apartmentPrefs'
import { useApartmentsStore } from '../../DAL/stores/apartments'
import { useAuthStore } from '../../DAL/stores/auth'
import { DataObjectState } from '../../DAL/stores/types'
import { App } from './App'

function AppContainer() {
  const user = useAuthStore(({ user }) => user)
  const getUserApartment = useApartmentsStore(({ getUserApartment }) => getUserApartment)
  const getUserApartmentPrefs = useApartmentPrefsStore(
    ({ getUserApartmentPrefs }) => getUserApartmentPrefs,
  )
  const apartment = useApartmentsStore(({ apartment }) => apartment)
  const apartmentPrefs = useApartmentPrefsStore(({ apartmentPrefs }) => apartmentPrefs)
  const navigate = useNavigate()
  const location = useLocation()

  useEffect(() => {
    if (user.data) {
      getUserApartment(user.data.id)
      getUserApartmentPrefs(user.data.id)
    }
    if (user.state === DataObjectState.Failed) {
      const from = location.state?.from?.pathname || '/'
      navigate('/login', { state: { from } })
    }
  }, [user])

  useEffect(() => {
    if (apartment.state === DataObjectState.Succeeded) {
      if (!apartment.data) {
        navigate('/apartment/add')
      } else if (apartmentPrefs.state === DataObjectState.Succeeded) {
        if (!apartmentPrefs.data || !apartmentPrefs.data.length) {
          navigate('/apartmentPref/add')
        } else {
          navigate('/')
        }
      }
    } else if (
      apartment.state === DataObjectState.Failed ||
      apartmentPrefs.state === DataObjectState.Failed
    ) {
      navigate('/error')
    }
  }, [apartment, apartmentPrefs])

  const isLoading =
    user.state === DataObjectState.Fetching ||
    user.state === DataObjectState.NotStarted ||
    apartment.state === DataObjectState.NotStarted ||
    apartment.state === DataObjectState.Fetching

  return <App user={user.data} isLoading={isLoading} />
}

export { AppContainer }
