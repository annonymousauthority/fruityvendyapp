<template>
  <div class="bg-gradient-to-br from-gray-900 to-gray-700 h-screen pt-12 lg:pt-64 p-6">
    <TransitionRoot as="template" :show="state.open">
      <Dialog as="div" class="relative z-10" @close="state.open = false">
        <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100"
          leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </TransitionChild>

        <div class="fixed inset-0 z-10 overflow-y-auto">
          <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <TransitionChild as="template" enter="ease-out duration-300"
              enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200"
              leave-from="opacity-100 translate-y-0 sm:scale-100"
              leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
              <DialogPanel
                class="relative transform overflow-hidden rounded-lg bg-white px-4 pt-5 pb-4 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-sm sm:p-6">
                <div>
                  <div class="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-green-100">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                      stroke="currentColor" class="w-6 h-6 text-green-500">
                      <path stroke-linecap="round" stroke-linejoin="round"
                        d="M15.75 15.75V18m-7.5-6.75h.008v.008H8.25v-.008zm0 2.25h.008v.008H8.25V13.5zm0 2.25h.008v.008H8.25v-.008zm0 2.25h.008v.008H8.25V18zm2.498-6.75h.007v.008h-.007v-.008zm0 2.25h.007v.008h-.007V13.5zm0 2.25h.007v.008h-.007v-.008zm0 2.25h.007v.008h-.007V18zm2.504-6.75h.008v.008h-.008v-.008zm0 2.25h.008v.008h-.008V13.5zm0 2.25h.008v.008h-.008v-.008zm0 2.25h.008v.008h-.008V18zm2.498-6.75h.008v.008h-.008v-.008zm0 2.25h.008v.008h-.008V13.5zM8.25 6h7.5v2.25h-7.5V6zM12 2.25c-1.892 0-3.758.11-5.593.322C5.307 2.7 4.5 3.65 4.5 4.757V19.5a2.25 2.25 0 002.25 2.25h10.5a2.25 2.25 0 002.25-2.25V4.757c0-1.108-.806-2.057-1.907-2.185A48.507 48.507 0 0012 2.25z" />
                    </svg>

                  </div>
                  <div class="mt-3 text-center sm:mt-5">
                    <DialogTitle as="h3" class="text-base font-semibold leading-6 text-gray-900">FruityVendyResult
                    </DialogTitle>
                    <div class="mt-2">
                      <p class="text-sm text-gray-500">Total Calories for your fruit punch: {{ state.calories }}</p>
                    </div>
                  </div>
                </div>
                <div class="mt-5 sm:mt-6">
                  <button type="button"
                    class="inline-flex w-full justify-center rounded-md border border-transparent bg-gradient-to-b from-green-500 to-green-800 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:text-sm"
                    @click="state.open = false">Ok</button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
    <div
      class="w-full lg:max-w-5xl lg:h-[650px] mx-auto flex flex-col shadow-xl shadow-gray-900 rounded-lg border-2 border-gray-800">
      <div class="block lg:flex justify-between items-center p-12">
        <h1 class="text-white text-2xl font-bold">FruityVendy</h1>
        <h2 class="text-white text-lg font-extralight italic">Made by Augustine</h2>
      </div>
      <div
        class="bg-gray-700 p-2 w-4/5 lg:w-1/3 text-white rounded-lg shadow-lg shadow-gray-800 mt-4 my-4 mx-auto text-sm flex items-center justify-between space-x-2">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="w-12 h-12">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M11.25 11.25l.041-.02a.75.75 0 011.063.852l-.708 2.836a.75.75 0 001.063.853l.041-.021M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9-3.75h.008v.008H12V8.25z" />
        </svg>
        <p>The calories unit is per 100gram of fruit, and also represent an approximate value for a unit of a fruit</p>
        <span></span>
      </div>
      <div
        class="w-full lg:w-1/3 mx-auto flex items-center justify-center bg-gray-800 shadow-lg shadow-gray-900 rounded-lg">
        <button @click="state.chat = true; state.dropdown = false"
          :class="[state.chat ? 'text-white rounded-l-lg bg-green-800 p-3 w-1/2 mx-0' : 'bg-transparent text-white w-1/2 p-3']">
          Chat with Vendy</button>
        <button @click="state.chat = false; state.dropdown = true"
          :class="[state.dropdown ? 'text-white rounded-r-lg bg-green-800 w-1/2 p-3' : 'bg-transparent text-white w-1/2 p-3']">Use
          Dropdown</button>
      </div>
      <div v-if="state.response != ''" class="my-3 w-full p-3 lg:p-0 lg:w-1/3 mx-auto flex text-white">
        <p> {{ state.response }}</p>
      </div>
      <div v-else></div>
      <Transition>
        <div v-if="state.dropdown" class="flex flex-col items-start p-12 pt-4">
          <div class="block lg:flex items-start w-full space-y-6 lg:space-y-0">
            <div class="w-full lg:w-1/3">
              <Combobox as="div" v-model="state.selectedFruits">
                <ComboboxLabel class="block text-sm font-medium text-white">Select Fruit</ComboboxLabel>
                <div class="relative mt-1">
                  <ComboboxInput
                    class="w-full rounded-md border border-gray-300 bg-white py-2 pl-3 pr-10 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500 sm:text-sm"
                    :display-value="(fruits) => state.selectedFruits === null ? 'Select Fruit' : fruits?.name" />
                  <ComboboxButton
                    class="absolute inset-y-0 right-0 flex items-center rounded-r-md px-2 focus:outline-none">
                    <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                  </ComboboxButton>
                  <ComboboxOptions v-if="filteredfruits.length > 0"
                    class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    <ComboboxOption v-for="(fruits, index) in filteredfruits" :key="index" :value="fruits" as="template"
                      v-slot="{ active, selected }">
                      <li
                        :class="['relative cursor-default select-none py-2 pl-3 pr-9', active ? 'bg-indigo-600 text-white' : 'text-gray-900']">
                        <div :class="['truncate flex justify-between space-x-6', selected && 'font-semibold']"
                          @click="onChangedSelection(fruits.name, index)">
                          <span @click="onChangedSelection(fruits.name, index)">{{ fruits.name }}</span>
                          <span>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                              stroke="currentColor" class="w-6 h-6">
                              <path stroke-linecap="round" stroke-linejoin="round"
                                d="M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                          </span>
                        </div>

                        <span v-if="selected"
                          :class="['absolute inset-y-0 right-0 flex items-center pr-4', active ? 'text-white' : 'text-indigo-600']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ComboboxOption>
                  </ComboboxOptions>
                </div>
              </Combobox>
            </div>
            <div class="w-full lg:w-1/3"></div>
            <div class="text-white w-full lg:w-1/3 flex items-start flex-col text-left space-y-6">
              <div class="flex justify-between space-x-3 items-center">
                <h1 class="text-2xl font-bold">Fruit Added</h1>
                <span :class="[maxFruits ? 'text-red-500' : 'text-green-500']">{{ state.fruits.length }}/5</span>
              </div>
              <TransitionGroup name="list" tag="ul" enter-active-class="transition-all duration-200 ease-in"
                leave-active-class="transition-all duration-500 ease-in" leave-to-class="opacity-0 translate-x-[30px]"
                enter-from-class="opacity-0 translate-x-[30px]">
                <li v-for="fruit in state.fruits" :key="fruit"
                  class="font-extralight flex space-x-6 justify-between mt-3">
                  <span>{{ fruit }}</span>
                  <button @click="remove(fruit)"><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                      stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-red-600">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 12h-15" />
                    </svg>
                  </button>
                  <button @click="increaseCalories()">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                      stroke="currentColor" class="w-6 h-6 text-green-500">
                      <path stroke-linecap="round" stroke-linejoin="round"
                        d="M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>

                  </button>
                </li>
              </TransitionGroup>
              <span class="text-sm font-extralight text-red-500" v-if="state.error">You don't atleast 3 fruits in your
                mix.</span>
              <span v-else></span>
              <button @click="showValue"
                class="mt-12 bg-gradient-to-b from-green-500 to-green-800 p-3 rounded-lg text-white shadow-lg shadow-gray-800">Calculate
                Calories</button>
            </div>
          </div>
        </div>
        <div v-else-if="state.chat" class="p-0 m-0">
          <div class="w-full lg:w-1/3 mx-auto flex items-center justify-center mt-3">
            <div>
              <label for="prompt" class="text-white text-sm">Enter Fruit Punch Prompt here.</label>
              <textarea id="prompt" v-model="state.prompt" :onchanged="onChangedReset" cols="7"
                class="w-full p-3 bg-white rounded-lg focus:ring-green-500 focus:outline-green-500 text-sm"> </textarea>
            </div>
          </div>
          <button @click="evaluatePrompt"
            class="mt-3 flex mx-auto bg-gradient-to-b from-green-500 to-green-800 p-3 rounded-lg text-white shadow-lg shadow-gray-800">Calculate
            Calories</button>
        </div>
      </Transition>

    </div>
  </div>
</template>


<script setup>
/*
- Run a goscript that will call all the fruits from api and update a json file with the fruits.
*/
import { computed, ref, reactive } from 'vue'
import { CheckIcon, ChevronUpDownIcon } from '@heroicons/vue/20/solid'
import {
  Combobox,
  ComboboxButton,
  ComboboxInput,
  ComboboxLabel,
  ComboboxOption,
  ComboboxOptions,
} from '@headlessui/vue'
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
import { fetchFruity } from './composables/Fetch'



const data = await fetchFruity("http://localhost:3000/getFruits") //change this to main servers when pushed to prod systems
const fruits = JSON.parse(JSON.stringify(data))

const state = reactive({
  fruits: [],
  query: '',
  selectedFruits: null,
  maxFruits: false,
  calories: 0,
  open: false,
  error: false,
  chat: false,
  dropdown: true,
  prompt: '',
  response: '',
  counter: 0,
  resLength: 0,
  res: ''
})

const filteredfruits = computed(() =>
  state.query === ''
    ? fruits
    : fruits.filter((fruits) => {
      return fruits.name.includes(state.query)
    })
)

const maxFruits = computed(() =>
  state.fruits.length >= 5
    ? true
    : false
)

function onChangedReset() {
  state.res = ''
  state.counter = 0
  state.response = ''
}

function updateFruits(value, id) {
  state.error = false
  if (!state.fruits.includes(value) && state.fruits.length < 5) {
    state.fruits.push(value)
    state.calories += fruits[id].nutritions.calories
  }
}

function evaluatePrompt() {
  const appName = 'Vendy'

  if (!state.prompt.includes(appName.toLowerCase)) {
    state.res = "My creators taught me it's polite to address someone by their name. My name is Vendy nice to meet you."
    setTimeout(() => {
      state.prompt = ''
      state.resLength = state.res.length
      typeResponse()
    }, 1200);
  }
}

function typeResponse() {
  if (state.counter < state.resLength) {
    state.response += state.res.charAt(state.counter)
    state.counter += 1
    setTimeout(typeResponse, 100);
  }
}

function onChangedSelection(value, id) {
  updateFruits(value, id)
}

function showValue() {
  if (state.fruits.length >= 3) {
    state.open = true
  } else {
    state.error = true
  }
}
function remove(fruit) {
  state.fruits.splice(state.fruits.indexOf(fruit), 1)
}


</script>