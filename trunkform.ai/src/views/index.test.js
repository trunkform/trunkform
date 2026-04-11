import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import IndexView from './index.vue'

describe('IndexView', () => {
  it('renders TrunkformHero and TrunkformTrademark components', () => {
    const wrapper = mount(IndexView)

    expect(wrapper.element.tagName)
      .toBe('MAIN')
    expect(wrapper.find('nav').exists())
      .toBe(true)
    expect(wrapper.find('nav').find('a').attributes('href'))
      .toBe('/')
    expect(wrapper.find('nav').find('a').text())
      .toBe('trunkform')
    expect(wrapper.findComponent({ name: 'TrunkformHero' }).exists())
      .toBe(true)
    expect(wrapper.findComponent({ name: 'TrunkformTrademark' }).exists())
      .toBe(true)
  })
})
