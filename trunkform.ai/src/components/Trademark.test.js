import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { readFileSync } from 'fs'
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'
import Trademark from './Trademark.vue'

const __dirname = dirname(fileURLToPath(import.meta.url))

describe('Trademark.vue', () => {
  it('renders correctly', () => {
    const wrapper = mount(Trademark)
    expect(wrapper.text()).toContain('trunkform is a trademark of Richard James Minchuk')
  })

  it('style block has not changed', () => {
    const src = readFileSync(join(__dirname, 'Trademark.vue'), 'utf-8')
    const style = src.match(/<style[^>]*>([\s\S]*?)<\/style>/)?.[1]
    expect(style).toMatchSnapshot()
  })
})
