package icu.aimerny.elix.data;

import icu.aimerny.elix.data.loot.ElixOreLootTable;
import icu.aimerny.elix.data.tag.ElixBlockTagGenerator;
import icu.aimerny.elix.registry.ModItem;
import net.fabricmc.fabric.api.datagen.v1.DataGeneratorEntrypoint;
import net.fabricmc.fabric.api.datagen.v1.FabricDataGenerator;
import net.fabricmc.fabric.api.datagen.v1.FabricDataOutput;
import net.fabricmc.fabric.api.datagen.v1.provider.FabricRecipeProvider;
import net.fabricmc.fabric.api.datagen.v1.provider.FabricTagProvider;
import net.minecraft.data.server.recipe.RecipeExporter;
import net.minecraft.data.server.recipe.RecipeProvider;
import net.minecraft.item.Item;
import net.minecraft.item.Items;
import net.minecraft.recipe.book.RecipeCategory;
import net.minecraft.registry.RegistryKeys;
import net.minecraft.registry.RegistryWrapper;
import net.minecraft.registry.tag.ItemTags;
import net.minecraft.registry.tag.TagKey;
import net.minecraft.util.Identifier;

import java.util.List;
import java.util.concurrent.CompletableFuture;


public class ElixDataGenerator implements DataGeneratorEntrypoint {

    @Override
    public void onInitializeDataGenerator(FabricDataGenerator fabricDataGenerator) {
        FabricDataGenerator.Pack pack = fabricDataGenerator.createPack();
        pack.addProvider(ElixItemTagGenerator::new);
        pack.addProvider(ElixBlockTagGenerator::new);
        pack.addProvider(ElixRecipeGenerator::new);
        pack.addProvider(ElixOreLootTable::new);
    }

    public static class ElixItemTagGenerator extends FabricTagProvider.ItemTagProvider{

        private static final TagKey<Item> SMELLY_ITEMS = TagKey.of(RegistryKeys.ITEM, new Identifier("elix:smelly_items"));

        public ElixItemTagGenerator(FabricDataOutput output, CompletableFuture<RegistryWrapper.WrapperLookup> completableFuture) {
            super(output, completableFuture);
        }

        @Override
        protected void configure(RegistryWrapper.WrapperLookup arg) {
            getOrCreateTagBuilder(SMELLY_ITEMS)
                    .add(Items.SLIME_BALL)
                    .add(Items.ROTTEN_FLESH)
                    .addOptionalTag(ItemTags.DIRT);
        }
    }

    public static class ElixRecipeGenerator extends FabricRecipeProvider {
        public ElixRecipeGenerator(FabricDataOutput output) {
            super(output);
        }

        @Override
        public void generate(RecipeExporter exporter) {
            RecipeProvider.offerSmelting(exporter, List.of(Items.OBSIDIAN),
                    RecipeCategory.BREWING, ModItem.OBSIDIAN_INGOT, 0.5F, 300, "elix");
        }
    }

}
