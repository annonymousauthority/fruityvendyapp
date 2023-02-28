<template>
  <div class="bg-gradient-to-br from-gray-900 to-gray-700 h-screen pt-64">
    <div
      class="w-full lg:max-w-5xl lg:h-[650px] mx-auto flex flex-col shadow-xl shadow-gray-900 rounded-lg border-2 border-gray-800">
      <div class="flex justify-between items-center p-12">
        <h1 class="text-white text-2xl font-bold">FruityVendy</h1>
        <h2 class="text-white text-lg font-extralight italic">Made by Augustine</h2>
      </div>
      <div class="flex flex-col items-start p-12">
        <div class="flex items-start w-full">
          <div class="w-1/3">
            <Combobox as="div" v-model="state.selectedFruits">
              <ComboboxLabel class="block text-sm font-medium text-white">Select Fruit</ComboboxLabel>
              <div class="relative mt-1">
                <ComboboxInput
                  class="w-full rounded-md border border-gray-300 bg-white py-2 pl-3 pr-10 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500 sm:text-sm"
                  :display-value="(person) => state.selectedFruits === null ? 'Select Fruit' : person?.name" />
                <ComboboxButton class="absolute inset-y-0 right-0 flex items-center rounded-r-md px-2 focus:outline-none">
                  <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                </ComboboxButton>

                <ComboboxOptions v-if="filteredPeople.length > 0"
                  class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                  <ComboboxOption v-for="person in filteredPeople" :key="person.id" :value="person" as="template"
                    v-slot="{ active, selected }">
                    <li
                      :class="['relative cursor-default select-none py-2 pl-3 pr-9', active ? 'bg-indigo-600 text-white' : 'text-gray-900']">
                      <div :class="['truncate flex justify-between space-x-6', selected && 'font-semibold']"
                        @click="onChangedSelection(person.name)">
                        <span>{{ person.name }}</span>
                        <span><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
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
            <button @click="showValue"
              class="mt-6 bg-gradient-to-b from-green-500 to-green-800 p-3 rounded-lg text-white shadow-lg shadow-gray-800">Show
              value</button>
          </div>
          <div class="w-1/3"></div>
          <div class="text-white w-1/3 flex items-start flex-col space-y-3 text-left">
            <h1 class="text-2xl font-bold">Fruit Added</h1>
            <TransitionGroup name="list" tag="ul" enter-active-class="transition-all duration-200 ease-in"
              leave-active-class="transition-all duration-500 ease-in" leave-to-class="opacity-0 translate-x-[30px]"
              enter-from-class="opacity-0 translate-x-[30px]">
              <li v-for="fruit in state.fruits" :key="fruit" class="font-extralight flex space-x-24 justify-between mt-3">
                <span>{{ fruit }}</span>
                <button @click="remove(fruit)"><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                    stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-red-600">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 12h-15" />
                  </svg>
                </button>
              </li>
            </TransitionGroup>
          </div>
        </div>



      </div>

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

import '~~/composables/Fetch'

const people = [
  { id: 1, name: 'Mango' },
  { id: 2, name: 'Strawberry' },
  { id: 2, name: 'Orange' },
]



const filteredPeople = computed(() =>
  state.query === ''
    ? people
    : people.filter((person) => {
      return person.name.includes(state.query)
    })
)

const state = reactive({
  fruits: [],
  query: '',
  selectedFruits: null,
})
function updateFruits(value) {
  if (!state.fruits.includes(value)) {
    state.fruits.push(value)
  }
}

function onChangedSelection(value) {
  updateFruits(value)
}

function showValue() {
  console.log(state.fruits)
}
function remove(fruit) {
  state.fruits.splice(state.fruits.indexOf(fruit), 1)
}


</script>