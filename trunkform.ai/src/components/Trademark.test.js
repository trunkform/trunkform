import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Trademark from './Trademark.vue'

describe('Trademark.vue', () => {
  it('renders correctly', () => {
    const wrapper = mount(Trademark)
    expect(wrapper.text()).toContain('trunkform is a trademark of Richard James Minchuk')
  })
})
