import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import IndexView from './index.vue'

describe('IndexView', () => {
  it('renders Hero and Trademark components', () => {
    const wrapper = mount(IndexView)

    expect(wrapper.findComponent({ name: 'Hero' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'Trademark' }).exists()).toBe(true)
  })
})
