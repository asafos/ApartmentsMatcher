import { ApartmentToAdd } from '../../../DAL/services/apartments/apartments'
import { Formik, Field, FieldProps } from 'formik'
import { Box, Flex, VStack, ButtonGroup } from '@chakra-ui/react'
import {
  CheckboxControl,
  FormControl,
  NumberInputControl,
  SelectControl,
  SubmitButton,
} from 'formik-chakra-ui'
import { RangeDatepicker } from 'chakra-dayzed-datepicker'
import * as Yup from 'yup'

type Props = {
  onSave: (apartment: Omit<ApartmentToAdd, 'user_id'>) => void
}

const today = new Date()
const yesterday = new Date()
yesterday.setDate(yesterday.getDate() - 1)
const oneYearFromToday = new Date()
oneYearFromToday.setFullYear(oneYearFromToday.getFullYear() + 1)

const validationSchema = Yup.object({
  numberOfRooms: Yup.number().min(1).max(30).required('number of rooms must be at least 1'),
  price: Yup.number().min(1).max(40000).required('price cannot be 0'),
  location: Yup.string().required('must choose a location'),
})

const AddApartmentPref = (props: Props) => {
  const { onSave } = props
  return (
    <Flex align='center' justify='center' h='100%'>
      <Box bg='white' p={6} rounded='md' w={64} minWidth='300px'>
        <Formik
          initialValues={{
            numberOfRooms: 0,
            price: 0,
            location: '',
            availableDates: [today, today],
            balcony: false,
            roof: false,
            parking: false,
            elevator: false,
            petsAllowed: false,
            renovated: false,
          }}
          validationSchema={validationSchema}
          onSubmit={(values) => {
            const formattedValues = {
              ...values,
              // availableDates: values.availableDates.map(d => d.toISOString()),
              price: Number(values.price),
              numberOfRooms: Number(values.numberOfRooms),
            }
            onSave(formattedValues)
          }}
        >
          {(formProps) => {
            const { handleSubmit, errors, touched } = formProps

            return (
              <VStack spacing={4} align='flex-start'>
                <NumberInputControl
                  isInvalid={!!errors.numberOfRooms && touched.numberOfRooms}
                  name='numberOfRooms'
                  label='Rooms'
                />
                <NumberInputControl
                  showStepper={false}
                  isInvalid={!!errors.price && touched.price}
                  name='price'
                  label='Price'
                />
                <SelectControl
                  name='location'
                  isInvalid={!!errors.location && touched.location}
                  selectProps={{ placeholder: 'Select location' }}
                >
                  <option value='LevTLV'>Lev TLV</option>
                  <option value='OldNorth'>Old North</option>
                  <option value='NewNorth'>New North</option>
                  <option value='Sarona'>Sarona</option>
                  <option value='NeveTzedek'>Neve Tzedek</option>
                  <option value='NeveShaanan'>Neve Shaanan</option>
                  <option value='Florentin'>Florentin</option>
                  <option value='RamatAviv'>Ramat Aviv</option>
                </SelectControl>
                <FormControl label='Available Dates' name='availableDates'>
                  <Field name={'availableDates'} id={'availableDates'}>
                    {({ field: { value }, form: { setFieldValue } }: FieldProps<Date[]>) => {
                      return (
                        value && (
                          <RangeDatepicker
                            selectedDates={value || []}
                            minDate={yesterday}
                            maxDate={oneYearFromToday}
                            onDateChange={(dates) => {
                              setFieldValue('availableDates', dates)
                            }}
                          />
                        )
                      )
                    }}
                  </Field>
                </FormControl>
                <Flex justifyContent='space-between' width='100%' mt='8px'>
                  <Flex direction='column'>
                    <CheckboxControl name='balcony' value='balcony'>
                      Balcony
                    </CheckboxControl>
                    <CheckboxControl name='roof' value='roof'>
                      Roof
                    </CheckboxControl>
                    <CheckboxControl name='parking' value='parking'>
                      Parking
                    </CheckboxControl>
                  </Flex>
                  <Flex direction='column'>
                    <CheckboxControl name='elevator' value='elevator'>
                      Elevator
                    </CheckboxControl>
                    <CheckboxControl name='petsAllowed' value='petsAllowed'>
                      Pets Allowed
                    </CheckboxControl>
                    <CheckboxControl name='renovated' value='renovated'>
                      Renovated
                    </CheckboxControl>
                  </Flex>
                </Flex>
                <ButtonGroup pt='16px' justifyContent='flex-end' width='100%'>
                  <SubmitButton onClick={() => handleSubmit()}>Submit</SubmitButton>
                  {/* <ResetButton>Reset</ResetButton> */}
                </ButtonGroup>
              </VStack>
            )
          }}
        </Formik>
      </Box>
    </Flex>
  )
}

export { AddApartmentPref }
