import React from 'react'
import ReactDOM from 'react-dom/client'
import { Main } from './Main'
import reportWebVitals from './reportWebVitals'
import { BrowserRouter } from 'react-router-dom'
import { ChakraProvider, extendTheme } from '@chakra-ui/react'
import { MultiSelectTheme } from 'chakra-multiselect'

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)

const theme = extendTheme({
  components: {
    MultiSelect: MultiSelectTheme,
  },
})

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <ChakraProvider theme={theme}>
        <Main />
      </ChakraProvider>
    </BrowserRouter>
  </React.StrictMode>,
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
