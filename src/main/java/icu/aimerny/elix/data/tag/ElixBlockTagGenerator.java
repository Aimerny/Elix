package icu.aimerny.elix.data.tag;

import icu.aimerny.elix.registry.ModBlock;
import net.fabricmc.fabric.api.datagen.v1.FabricDataOutput;
import net.fabricmc.fabric.api.datagen.v1.provider.FabricTagProvider;
import net.minecraft.block.Block;
import net.minecraft.item.Item;
import net.minecraft.registry.RegistryKeys;
import net.minecraft.registry.RegistryWrapper;
import net.minecraft.registry.tag.TagKey;
import net.minecraft.util.Identifier;

import java.util.concurrent.CompletableFuture;

public class ElixBlockTagGenerator extends FabricTagProvider.BlockTagProvider {
    public ElixBlockTagGenerator(FabricDataOutput output, CompletableFuture<RegistryWrapper.WrapperLookup> registriesFuture) {
        super(output, registriesFuture);
    }
    //private static final TagKey<Item> SMELLY_ITEMS = TagKey.of(RegistryKeys.ITEM, new Identifier("elix:smelly_items"));
    private static final TagKey<Block> NEEDS_STONE_TOOL = TagKey.of(RegistryKeys.BLOCK, new Identifier("minecraft:need_stone_tool"));
    private static final TagKey<Block> PICKAXE = TagKey.of(RegistryKeys.BLOCK, new Identifier("minecraft:mineable/pickaxe"));

    @Override
    protected void configure(RegistryWrapper.WrapperLookup arg) {
        getOrCreateTagBuilder(PICKAXE)
                .add(ModBlock.OBSIDIAN_ORE);

        getOrCreateTagBuilder(NEEDS_STONE_TOOL)
                .add(ModBlock.OBSIDIAN_ORE);
    }
}
