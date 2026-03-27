import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { readFileSync } from 'fs'
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'
import Hero from './Hero.vue'

const __dirname = dirname(fileURLToPath(import.meta.url))

describe('Hero.vue', () => {
  it('renders correctly', () => {
    const wrapper = mount(Hero)
    expect(wrapper.text()).toContain('trunkform')
  })

  it('style block has not changed', () => {
    const src = readFileSync(join(__dirname, 'Hero.vue'), 'utf-8')
    const style = src.match(/<style[^>]*>([\s\S]*?)<\/style>/)?.[1]
    expect(style).toMatchSnapshot()
  })
})
