import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { readFileSync } from 'fs'
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'
import TrunkformHero from './TrunkformHero.vue'

const __dirname = dirname(fileURLToPath(import.meta.url))

describe('TrunkformHero.vue', () => {
  it('renders correctly', () => {
    const wrapper = mount(TrunkformHero)
    expect(wrapper.find('img').attributes('src'))
      .toBe('/hero.png')
    expect(wrapper.find('img').attributes('alt'))
      .toBe('')

    expect(wrapper.find('section').exists())
      .toBe(true)
    expect(wrapper.find('section').find('h1').text())
      .toBe('Designed for Speed, Built for Trust')
    expect(wrapper.find('section').find('p').text())
      .toBe('trunkform™ MCP helps developers remove process bottlenecks, standardize delivery, and build release ready software.')
    expect(wrapper.find('section').find('div').find('a:first-child').attributes('href'))
      .toBe('https://github.com/trunkform/trunkform/blob/trunk/mcp/README.md')
    expect(wrapper.find('section').find('div').find('a:first-child').text())
      .toBe('learn more')
    expect(wrapper.find('section').find('div').find('a:last-child').attributes('href'))
      .toBe('https://github.com/trunkform/trunkform')
    expect(wrapper.find('section').find('div').find('a:last-child').text())
      .toBe('view the source')
  })

  it('style block has not changed', () => {
    const src = readFileSync(join(__dirname, 'TrunkformHero.vue'), 'utf-8')
    const style = src.match(/<style[^>]*>([\s\S]*?)<\/style>/)?.[1]
    expect(style).toMatchSnapshot()
  })
})
