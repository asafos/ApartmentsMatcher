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
import { ApartmentPrefToAdd } from '../../../DAL/services/apartmentPrefs/apartmentPrefs'
import { locationLabelsMap } from '../../../DAL/services/types'
import { MultiSelect } from 'chakra-multiselect'

type Props = {
  onSave: (apartment: Omit<ApartmentPrefToAdd, 'user_id'>) => void
}

const today = new Date()
const yesterday = new Date()
yesterday.setDate(yesterday.getDate() - 1)
const oneYearFromToday = new Date()
oneYearFromToday.setFullYear(oneYearFromToday.getFullYear() + 1)

const validationSchema = Yup.object({
  numberOfRoomsMin: Yup.number().min(1).max(30).required('number of rooms must be at least 1'),
  numberOfRoomsMax: Yup.number().min(1).max(30).required('number of rooms must be at least 1'),
  priceMin: Yup.number().min(1).max(40000).required('price cannot be 0'),
  priceMax: Yup.number().min(1).max(40000).required('price cannot be 0'),
  location: Yup.array().of(Yup.string()).min(1).required('must choose a location'),
})

const AddApartmentPref = (props: Props) => {
  const { onSave } = props
  return (
    <Flex align='center' justify='center' h='100%'>
      <Box bg='white' p={6} rounded='md' w={64} minWidth='340px'>
        <Formik
          initialValues={
            {
              numberOfRoomsMin: 1,
              numberOfRoomsMax: 2,
              priceMin: 1000,
              priceMax: 2000,
              location: [],
              availableDates: [today, today],
              balcony: false,
              roof: false,
              parking: false,
              elevator: false,
              petsAllowed: false,
              renovated: false,
            } as Omit<ApartmentPrefToAdd, 'user_id' | 'numberOfRooms' | 'price'> & {
              numberOfRoomsMin: number
              numberOfRoomsMax: number
              priceMin: number
              priceMax: number
            }
          }
          validate={(values) => {
            const errors: any = {}
            if (values.priceMin > values.priceMax) {
              errors.priceMax = 'Maximum price cannot go below maximum price'
            } else if (values.numberOfRoomsMin > values.numberOfRoomsMax) {
              errors.numberOfRoomsMax = 'Maximum rooms cannot go below maximum rooms'
            }
            return errors
          }}
          validationSchema={validationSchema}
          onSubmit={({ priceMin, priceMax, numberOfRoomsMin, numberOfRoomsMax, ...values }) => {
            const formattedValues: Omit<ApartmentPrefToAdd, 'user_id'> = {
              ...values,
              // availableDates: values.availableDates.map(d => d.toISOString()),
              price: [Number(priceMin), Number(priceMax)],
              numberOfRooms: [Number(numberOfRoomsMin), Number(numberOfRoomsMax)],
            }
            onSave(formattedValues)
          }}
        >
          {(formProps) => {
            const { handleSubmit, errors, touched, setSubmitting } = formProps

            return (
              <VStack spacing={4} align='flex-start'>
                <NumberInputControl
                  isInvalid={!!errors.numberOfRoomsMin && touched.numberOfRoomsMin}
                  name='numberOfRoomsMin'
                  label='Minimum Rooms'
                />

                <NumberInputControl
                  isInvalid={!!errors.numberOfRoomsMax && touched.numberOfRoomsMax}
                  name='numberOfRoomsMax'
                  label='Maximum Rooms'
                />

                <NumberInputControl
                  isInvalid={!!errors.priceMin && touched.priceMin}
                  name='priceMin'
                  label='Minimum Price'
                />

                <NumberInputControl
                  isInvalid={!!errors.priceMax && touched.priceMax}
                  name='priceMax'
                  label='Maximum Price'
                />

                <FormControl label='Locations' name='location'>
                  <Field name={'location'} id={'lcoation'}>
                    {({ field: { value }, form: { setFieldValue } }: FieldProps<Date[]>) => {
                      console.log(
                        '!@# ~ file: AddApartmentPref.tsx:113 ~ AddApartmentPref ~ value:',
                        value,
                      )

                      return (
                        value && (
                          <MultiSelect
                            options={Object.entries(locationLabelsMap).map(([value, label]) => ({
                              label,
                              value,
                            }))}
                            value={value}
                            label='Select location'
                            onChange={(val: any) => {
                              setFieldValue(
                                'location',
                                val.map((value: any) =>
                                  typeof value === 'string' ? value : value.value,
                                ),
                              )
                            }}
                          />
                        )
                      )
                    }}
                  </Field>
                </FormControl>

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
